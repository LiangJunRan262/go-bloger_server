package models

import "bloger_server/models/enum"

type LogModel struct {
	Model
	LogType     enum.LogType      `json:"logType"`                    // 日志类型 1:登录 2:操作 3
	Title       string            `gorm:"size:256" json:"title"`      // 标题
	Content     string            `json:"content"`                    // 内容
	Level       enum.LogLevelType `json:"level"`                      // 日志等级 1:普通 2:警告 3:错误
	UserID      uint              `json:"userID"`                     // 用户id
	UserModel   UserModel         `gorm:"foreignKey:UserID" json:"-"` // 用户信息
	IP          string            `gorm:"size:32" json:"ip"`          // IP地址
	Address     string            `gorm:"size:128" json:"address"`    // 地址
	IsRead      bool              `json:"isRead"`                     // 是否已读
	LoginStatus int8              `json:"loginStatus"`                // 登录状态 1:成功 2:失败
	Username    string            `gorm:"size:32" json:"username"`    // 用户名
	Password    string            `gorm:"size:32" json:"password"`    // 密码 (脱敏处理)
	LoginType   enum.LoginType    `json:"loginType"`                  // 登录类型 1:账号密码登录 2:验证码登录 3:第三方登录
}
