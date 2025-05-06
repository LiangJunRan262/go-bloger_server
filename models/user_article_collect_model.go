package models

import "time"

type UserArticleCollectModel struct {
	UserID       uint         `gorm:"uniqueIndex:idx_name" json:"userID"`    // 用户id
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`            // 用户信息
	ArticleID    uint         `gorm:"uniqueIndex:idx_name" json:"articleID"` // 文章id
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`         // 文章信息
	CollectID    uint         `gorm:"uniqueIndex:idx_name" json:"collectID"` // 收藏夹id
	CollectModel CollectModel `gorm:"foreignKey:CollectID" json:"-"`         // 收藏夹信息
	CreatedAt    time.Time    `gorm:"autoCreateTime" json:"createdAt"`       // 收藏时间
}
