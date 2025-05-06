package router

import (
	"bloger_server/global"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	fmt.Println("addr", global.Config.System.GetAddr())

	// 启动web服务
	r := gin.Default()

	fmt.Println("addr", global.Config.System.GetAddr())

	// 注册路由
	nr := r.Group("/api")

	SiteRouter(nr)

	addr := global.Config.System.GetAddr()
	fmt.Println("addr", addr)
	r.Run(addr)
}
