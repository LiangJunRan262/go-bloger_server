package common

import (
	"bloger_server/global"
	"fmt"
	"gorm.io/gorm"
)

type PageInfo struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Order string `form:"order"` // 可以覆盖
}

func (p PageInfo) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

func (p PageInfo) GetLimit() int {
	if p.Limit > 100 || p.Limit <= 0 {
		return 10
	}
	return p.Limit
}

func (p PageInfo) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

type Option struct {
	PageInfo     PageInfo
	Likes        []string
	Preloads     []string
	Where        *gorm.DB
	Debug        bool
	DefaultOrder string `form:"order"`
}

func ListQuery[T any](model T, option Option) (list []T, count int, err error) {
	// 自己的基础查询
	query := global.DB.Model(model).Where(model)

	// 日志
	if option.Debug {
		query = query.Debug()
	}

	// 模糊匹配
	if len(option.Likes) > 0 && option.PageInfo.Key != "" {
		likes := global.DB.Where("")
		for _, column := range option.Likes {
			likes.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.PageInfo.Key))
		}
		query = query.Where(likes)
	}

	// 自定义查询
	if option.Where != nil {
		query = query.Where(option.Where)
	}

	// 预加载
	if len(option.Preloads) > 0 {
		for _, column := range option.Preloads {
			query = query.Preload(column)
		}
	}

	// 查总数
	var _count int64
	query.Count(&_count)
	count = int(_count)

	// 分页
	limit := option.PageInfo.GetLimit()
	offset := option.PageInfo.GetOffset()

	// 排序
	if option.PageInfo.Order != "" {
		query = query.Order(option.PageInfo.Order)
	} else {
		if option.DefaultOrder != "" {
			query = query.Order(option.DefaultOrder)
		}
	}

	err = query.Debug().Offset(offset).Limit(limit).Find(&list).Error
	return list, count, err
}
