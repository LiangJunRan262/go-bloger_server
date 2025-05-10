package models

type CommentModel struct {
	Model
	UserID         uint            `json:"userID"`                        // 用户id
	UserModel      UserModel       `gorm:"foreignKey:UserID" json:"-"`    // 用户信息
	ArticleID      uint            `json:"articleID"`                     // 文章id
	ArticleModel   ArticleModel    `gorm:"foreignKey:ArticleID" json:"-"` // 文章信息
	Content        string          `gorm:"size:256" json:"content"`       // 评论内容
	ParentID       *uint           `json:"parentID"`                      // 父评论id
	ParentModel    *CommentModel   `gorm:"foreignKey:ParentID" json:"-"`  // 父评论信息
	SubCommentList []*CommentModel `gorm:"foreignKey:ParentID" json:"-"`  // 子评论列表
	RootParentID   *uint           `json:"rootParentID"`                  // 根评论id
	LikeCount      int             `json:"likeCount"`                     // 点赞数量
}
