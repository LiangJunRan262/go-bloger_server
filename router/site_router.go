package router

import (
	"bloger_server/api"

	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.SiteApi
	// 注册路由
	r.GET("site", app.SiteInfoView)
	//
	r.PUT("site/update", app.SiteUpdateView)
}
