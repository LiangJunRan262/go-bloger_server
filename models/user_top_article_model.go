package models

import "time"

type UserTopArticleModel struct {
	UserID       uint         `gorm:"uniqueIndex:idx_name" json:"userID"`    // 用户id
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`            // 用户信息
	ArticleID    uint         `gorm:"uniqueIndex:idx_name" json:"articleID"` // 文章id
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`         // 文章信息
	CreatedAt    time.Time    `gorm:"autoCreateTime" json:"createdAt"`       // 置顶时间
}
