package core

import (
	"bloger_server/global"
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	r := global.Config.Redis

	redisDB := redis.NewClient(&redis.Options{
		Addr:     r.Addr,     // 不写默认就是这个
		Username: r.Username, // 用户名
		Password: r.Password, // 密码
		DB:       r.DB,       // 默认是0
	})
	_, err := redisDB.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatalf("redis链接失败: %s", err.Error())
	}
	logrus.Info("redis链接成功")
	return redisDB
}
