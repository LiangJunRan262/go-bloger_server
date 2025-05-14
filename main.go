package main

import (
	"bloger_server/core"
	"bloger_server/flags"
	"bloger_server/global"
	"bloger_server/router"

	_ "github.com/sirupsen/logrus"
)

func main() {
	// 初始化配置
	flags.Parse()
	global.Config = core.ReadConf()
	global.DB = core.InitDB()
	global.Redis = core.InitRedis()
	core.InitLogrus()

	// 初始化数
	flags.Run()

	// 启动web服务
	router.Run()
}
