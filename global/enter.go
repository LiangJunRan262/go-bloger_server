package global

import (
	"bloger_server/conf"

	"gorm.io/gorm"
)

var Config *conf.Config

var DB *gorm.DB
