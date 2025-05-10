package models

type ArticleModel struct {
	Model
	Title         string    `gorm:"size:256" json:"title"`                        // 标题
	Abstract      string    `gorm:"size:256" json:"abstract"`                     // 简介
	Content       string    `gorm:"type:longtext" json:"content"`                 // 内容
	CategoryID    uint      `json:"categoryID"`                                   // 分类id
	TagList       []string  `gorm:"type:longtext;serializer:json" json:"tagList"` // 标签
	Cover         string    `gorm:"size:256" json:"cover"`                        // 封面
	UserID        uint      `json:"userID"`                                       // 用户id
	UserModel     UserModel `gorm:"foreignKey:UserID" json:"-"`                   // 用户信息
	LookCount     int       `json:"lookCount"`                                    // 浏览量
	LikeCount     int       `json:"likeCount"`                                    // 点赞量
	CommentCount  int       `json:"commentCount"`                                 // 评论量
	CoollectCount int       `json:"coollectCount"`                                // 收藏量
	OpenComment   bool      `json:"openComment"`                                  // 开放评论
	Status        int8      `json:"status"`                                       // 状态 0: 草稿 1: 审核中 2: 已经发布 3: 已删除 4: 已下架
}
