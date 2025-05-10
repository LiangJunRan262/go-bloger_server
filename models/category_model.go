package models

type CategoryModel struct {
	Model
	Title     string    `gorm:"size:32" json:"title"`       // 分类名称
	UserID    uint      `json:"userID"`                     // 用户id
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"` // 用户信息
}
