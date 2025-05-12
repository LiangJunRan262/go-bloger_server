package router

import (
	"bloger_server/api"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi

	r.GET("/log_list", app.LogListView)
}
