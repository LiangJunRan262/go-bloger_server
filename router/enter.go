package router

import (
	"bloger_server/global"
	"bloger_server/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(global.Config.System.GinMode)

	// 启动web服务
	r := gin.Default()

	// 注册静态文件
	r.Static("/static", "./static")

	// 注册路由
	nr := r.Group("/api")

	// 注册中间件
	nr.Use(middleware.LogMiddleware)

	SiteRouter(nr)

	addr := global.Config.System.GetAddr()
	r.Run(addr)
}
