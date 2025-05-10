package models

type CollectModel struct {
	Model
	Title        string    `gorm:"size:32" json:"title"`       // 收藏标题
	Abstract     string    `gorm:"size:32" json:"abstract"`    // 收藏摘要
	Cover        string    `gorm:"size:32" json:"cover"`       // 收藏封面
	ArticleCount int       `json:"articleCount"`               // 收藏文章数量
	UserID       uint      `json:"userID"`                     // 用户id
	UserModel    UserModel `gorm:"foreignKey:UserID" json:"-"` // 用户信息
}
