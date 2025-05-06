package models

// 全局通知模型
type GlobalNotificationModel struct {
	Model
	Title   string `grom:"size:64" json:"title"` // 标题
	Content string `json:"content"`              // 内容
	Icon    string `grom:"size:64" json:"icon"`  // 图标url
	Url     string `grom:"size:256" json:"url"`  // 跳转url
	Status  int    `json:"status"`               // 状态 0:未发布 1:已发布 2:已删除
}

// 用户读取全局通知消息表
type UserReadGlobalNotificationModel struct {
	Model
	UserID            uint                    `gorm:"uniqueIndex:idx_name" json:"userID"`         // 用户id
	UserModel         UserModel               `gorm:"foreignKey:UserID" json:"-"`                 // 用户信息
	NotificationID    uint                    `gorm:"uniqueIndex:idx_name" json:"notificationID"` // 通知id
	NotificationModel GlobalNotificationModel `gorm:"foreignKey:NotificationID" json:"-"`         // 通知信息
}
