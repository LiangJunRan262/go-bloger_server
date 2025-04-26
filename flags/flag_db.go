package flags

import (
	"bloger_server/global"
	"bloger_server/models"

	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},
	)

	if err != nil {
		logrus.Errorf("数据库迁移失败: %s", err.Error())
		return
	}
	logrus.Info("数据库迁移成功")
}
