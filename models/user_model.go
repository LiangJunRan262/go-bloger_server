package models

type UserModel struct {
	Model
	Username       string `grom:"size:32" json:"username"`
	Nickname       string `grom:"size:32" json:"nickname"`
	Avatar         string `grom:"size:256" json:"avatar"`
	Abstract       string `gorm:"size:256" json:"abstract"` // 简介
	RegisterSource int8   `json:"registerSource"`           // 注册来源
	CodeAge        int    `json:"codeAge"`                  // 码龄
	Password       string `grom:"size:64" json:"-"`
	// LikeTags       []string `gorm:"type:longtext;serializer:json" json:"likeTags"` // 喜欢的标签
	Email  string `grom:"size:256" json:"email"`
	Openid string `grom:"size:64" json:"openid"` // 第三发登录的openid
}
