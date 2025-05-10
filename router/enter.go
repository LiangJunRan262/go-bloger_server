package router

import (
	"bloger_server/global"
	"bloger_server/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	// 启动web服务
	r := gin.Default()

	// 注册路由
	nr := r.Group("/api")

	// 注册中间件
	nr.Use(middleware.LogMiddleware)

	SiteRouter(nr)

	addr := global.Config.System.GetAddr()
	fmt.Println("addr", addr)
	r.Run(addr)
}
