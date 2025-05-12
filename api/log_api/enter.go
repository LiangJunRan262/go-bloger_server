package log_api

import (
	"bloger_server/common/res"
	"bloger_server/global"
	"bloger_server/models"
	"bloger_server/models/enum"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	Limit     int               `form:"limit"`
	Page      int               `form:"page"`
	Key       string            `form:"key"`
	LogType   enum.LogType      `form:"logType"`
	Level     enum.LogLevelType `form:"level"`
	UserID    uint              `form:"userID"`
	IP        string            `form:"ip"`
	LoginType enum.LoginType    `form:"loginType"`
	Service   string            `form:"service"`
}

type LogListResponse struct {
	models.LogModel
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
}

func (LogApi) LogListView(c *gin.Context) {
	// 分页 查询（精确查询，模糊匹配）
	var cr LogListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWidthError(err, c)
		return
	}

	fmt.Printf("LogListView cr=%+v\n", cr)

	var list []models.LogModel
	if cr.Page <= 0 {
		cr.Page = 1
	}
	if cr.Page > 100 || cr.Limit <= 0 {
		cr.Page = 100
	}
	offset := (cr.Page - 1) * cr.Limit

	fmt.Printf("LogListView offset=%d limit=%d\n", offset, cr.Limit)

	model := models.LogModel{
		LogType:   cr.LogType,
		Level:     cr.Level,
		UserID:    cr.UserID,
		IP:        cr.IP,
		LoginType: cr.LoginType,
		Service:   cr.Service,
	}

	like := global.DB.Debug().Where("title like ?", fmt.Sprintf("%%%s%%", cr.Key))

	global.DB.Preload("UserModel").Debug().Where(like).Where(model).Offset(offset).Limit(cr.Limit).Find(&list)

	var count int64
	global.DB.Debug().Where(like).Where(model).Model(&models.LogModel{}).Count(&count)

	var _list = make([]LogListResponse, 0)
	for _, item := range list {
		_list = append(_list, LogListResponse{
			LogModel:     item,
			UserNickname: item.UserModel.Nickname,
			UserAvatar:   item.UserModel.Avatar,
		})
	}

	res.OkWithList(_list, int(count), c)
}
