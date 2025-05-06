package models

type UserArticleLookHistoryModel struct {
	Model
	UserID       uint         `json:"userID"`                        // 用户id
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`    // 用户信息
	ArticleID    uint         `json:"articleID"`                     // 文章id
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"` // 文章信息
}
