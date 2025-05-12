package log_api

import (
	"bloger_server/common"
	"bloger_server/common/res"
	"bloger_server/global"
	"bloger_server/models"
	"bloger_server/models/enum"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	common.PageInfo
	Key       string            `form:"key"`
	LogType   enum.LogType      `form:"logType"`
	Level     enum.LogLevelType `form:"level"`
	UserID    uint              `form:"userID"`
	IP        string            `form:"ip"`
	LoginType enum.LoginType    `form:"loginType"`
	Service   string            `form:"service"`
}

type LogListResponse struct {
	models.LogModel
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
}

func (LogApi) LogListView(c *gin.Context) {
	// 分页 查询（精确查询，模糊匹配）
	var cr LogListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWidthError(err, c)
		return
	}

	cr.PageInfo.Order = "id desc"
	list, count, err := common.ListQuery(models.LogModel{
		LogType:   cr.LogType,
		Level:     cr.Level,
		UserID:    cr.UserID,
		IP:        cr.IP,
		LoginType: cr.LoginType,
		Service:   cr.Service,
	}, common.Option{
		PageInfo: cr.PageInfo,
		Likes:    []string{"title"},
		Preloads: []string{"UserModel"},
		Debug:    true,
	})

	if err != nil {
		res.FailWithMsg("log日志查询失败", c)
	}

	var _list = make([]LogListResponse, 0)
	for _, item := range list {
		_list = append(_list, LogListResponse{
			LogModel:     item,
			UserNickname: item.UserModel.Nickname,
			UserAvatar:   item.UserModel.Avatar,
		})
	}

	res.OkWithList(_list, int(count), c)
}

func (LogApi) LogReadView(c *gin.Context) {
	var cr models.IDRequest

	if err := c.ShouldBindUri(&cr); err != nil {
		res.FailWidthError(err, c)
		return
	}

	var log models.LogModel
	err := global.DB.Debug().Take(&log, cr.ID).Error
	if err != nil {
		res.FailWithMsg("日志不存在", c)
		return
	}

	if !log.IsRead {
		global.DB.Model(&log).Update("is_read", true)
	}

	res.OkWithMsg("日志读取成功", c)
}
