package flags

import (
	"bloger_server/global"
	"bloger_server/models"

	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},                       // 用户表
		&models.UserConfModel{},                   // 用户配置表
		&models.ArticleModel{},                    // 文章表
		&models.CategoryModel{},                   // 分类表
		&models.ArticleLikeModel{},                // 文章点赞表
		&models.CollectModel{},                    // 收藏夹表
		&models.UserArticleCollectModel{},         // 用户文章收藏夹表
		&models.ImageModel{},                      // 图片表
		&models.UserArticleLookHistoryModel{},     // 用户文章浏览历史表
		&models.UserTopArticleModel{},             // 用户文章置顶表
		&models.CommentModel{},                    // 评论表
		&models.BannerModel{},                     // 轮播图表
		&models.LogModel{},                        // 日志表
		&models.UserLoginModel{},                  // 用户登录表
		&models.GlobalNotificationModel{},         // 全局通知表
		&models.UserReadGlobalNotificationModel{}, // 用户通知表
	)

	if err != nil {
		logrus.Fatalf("数据库迁移失败: %s", err.Error())
		return
	}
	logrus.Info("数据库迁移成功")
}
