package site_api

import (
	"bloger_server/models/enum"
	log_service "bloger_server/service/log_servive"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	log := log_service.GetActionLog(c)
	log.SetShowRequest()
	log.SetShowRequestHeader()
	log.SetShowResponse()
	log.SetShowResponseHeader()
	log.SetTitle("更新站点信息")
	log.SetItemInfo("请求时间", time.Now().Format("2006-01-02 15:04:05"))

	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		fmt.Println(err.Error())
		logrus.Errorf("参数错误: %v", err.Error())
		log.SetItemErr("错误信息", err)
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}

	log.SetItemInfo("cr", cr)
	//log.SetItemErr("错误信息", 123123321)

	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
