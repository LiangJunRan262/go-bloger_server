package models

type UserLoginModel struct {
	Model
	UserID    uint      `json:"userID"`                     // 用户id
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"` // 用户信息
	IP        string    `gorm:"size:32" json:"ip"`          // IP地址
	Address   string    `gorm:"size:128" json:"address"`    // 地址
	UA        string    `gorm:"size:128" json:"ua"`         // 设备信息
}
