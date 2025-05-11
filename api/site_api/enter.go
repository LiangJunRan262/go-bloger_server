package site_api

import (
	"bloger_server/common/res"
	log_service "bloger_server/service/log_servive"
	"time"

	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	res.Ok("data", "成功", c)
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
		log.SetItemErr("错误信息", err)
		res.FailWidthError(err, c)
		return
	}

	log.SetItemInfo("cr", cr)
	res.OkWithMsg("更新成功", c)
	return
}
