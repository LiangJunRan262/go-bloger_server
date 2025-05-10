package core

import (
	"bloger_server/global"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func InitDB() *gorm.DB {

	dc := global.Config.DB // 数据库配置

	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成外键约束
		SkipDefaultTransaction:                   true,
	})

	if err != nil {
		logrus.Fatal("连接数据库失败, error=" + err.Error())
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatal("获取数据库连接失败, error=" + err.Error())
		return nil
	}

	sqlDB.SetMaxIdleConns(dc.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dc.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	logrus.Info("数据库连接成功")

	err = db.Use(dbresolver.Register(dbresolver.Config{
		// use `db2` as sources, `db3`, `db4` as replicas
		Sources:  []gorm.Dialector{mysql.Open(dc.DSN())}, // 写的
		Replicas: []gorm.Dialector{mysql.Open(dc.DSN())},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
	}))

	if err != nil {
		logrus.Fatal("配置数据库读写分离失败, error=" + err.Error())
	}

	return db
}
