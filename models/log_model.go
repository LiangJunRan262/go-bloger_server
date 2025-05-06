package models

type LogModel struct {
	Model
	LogType   int8      `json:"logType"`                    // 日志类型 1:登录 2:操作 3
	Title     string    `gorm:"size:256" json:"title"`      // 标题
	Content   string    `gorm:"size:2048" json:"content"`   // 内容
	Level     int8      `json:"level"`                      // 日志等级 1:普通 2:警告 3:错误
	UserID    uint      `json:"userID"`                     // 用户id
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"` // 用户信息
	IP        string    `gorm:"size:32" json:"ip"`          // IP地址
	Address   string    `gorm:"size:128" json:"address"`    // 地址
	IsRead    bool      `json:"isRead"`                     // 是否已读

}
