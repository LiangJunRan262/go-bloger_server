package router

import (
	"bloger_server/api"
	"bloger_server/middleware"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi

	r.Use(middleware.AuthMiddleware)

	// 日志列表
	r.GET("/log_list", app.LogListView)
	// 读取日志
	r.GET("/log_read/:id", app.LogReadView)
	// 删除日志
	r.POST("/log_delete", app.LogDeleteView)
}
