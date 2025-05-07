package site_api

import (
	"bloger_server/models/enum"
	log_service "bloger_server/service/log_servive"

	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	log_service.NewLoginSuccess(c, enum.UserPwdLoginType)
	log_service.NewLoginFail(c, enum.UserPwdLoginType, "登录失败", "admin", "123456")
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
