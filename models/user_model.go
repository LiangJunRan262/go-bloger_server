package models

import "time"

type UserModel struct {
	Model
	Username       string `grom:"size:32" json:"username"`
	Nickname       string `grom:"size:32" json:"nickname"`
	Avatar         string `grom:"size:256" json:"avatar"`
	Abstract       string `gorm:"size:256" json:"abstract"` // 简介
	RegisterSource int8   `json:"registerSource"`           // 注册来源
	CodeAge        int    `json:"codeAge"`                  // 码龄
	Password       string `grom:"size:64" json:"-"`
	Email          string `grom:"size:256" json:"email"`
	Openid         string `grom:"size:64" json:"openid"` // 第三发登录的openid
	Role           int8   `json:"role"`                  // 角色 0:普通用户 1:管理员 3:访客
}

type UserConfModel struct {
	// Model
	UserID             uint       `gorm:"unique" json:"userID"` // 用户id
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"` // 兴趣标签
	UpdateUsernameDate *time.Time `json:"updateUsernameDate"`                            // 上次修改用户名的时间
	OpenCollect        bool       `json:"openCollect"`                                   // 公开我的收藏
	OpenFans           bool       `json:"openFans"`                                      // 公开我的粉丝
	OpenFollow         bool       `json:"openFollow"`                                    // 公开我的关注
	HomeStyleID        uint       `json:"homeStyleID"`                                   // 主页样式id
}
