package site_api

import (
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
