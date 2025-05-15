package global

import (
	"bloger_server/conf"
	"context"

	"github.com/go-redis/redis/v8"

	"gorm.io/gorm"
)

var Config *conf.Config

var DB *gorm.DB

var Redis *redis.Client

var CommonContext = context.Background()
