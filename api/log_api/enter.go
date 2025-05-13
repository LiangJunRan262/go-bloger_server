package log_api

import (
	"bloger_server/common"
	"bloger_server/common/res"
	"bloger_server/global"
	"bloger_server/models"
	"bloger_server/models/enum"
	log_service "bloger_server/service/log_servive"
	"fmt"
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

// 删除log
func (LogApi) LogDeleteView(c *gin.Context) {
	var cr models.DeleteRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWidthError(err, c)
		fmt.Println(err)
		return
	}

	log := log_service.GetActionLog(c)
	log.SetTitle("删除日志")
	log.SetShowResponse()
	log.SetShowRequest()

	fmt.Println(cr.IDs)

	var logList []models.LogModel
	err := global.DB.Debug().Find(&logList, "id in ?", cr.IDs).Error
	if err != nil {
		res.FailWithMsg("查询日志失败", c)
	}
	if len(logList) != len(cr.IDs) {
		// 列出不存在的日志ID
		var notExistIDs []uint
		for _, id := range cr.IDs {
			var exists bool
			for _, log := range logList {
				if log.ID == id {
					exists = true
					break
				}
			}
			if !exists {
				notExistIDs = append(notExistIDs, id)
			}
		}
		res.FailWithMsg(fmt.Sprintf("日志ID %v 不存在", notExistIDs), c)
		log.SetItemErr("日志ID不存在", fmt.Sprintf("日志ID %v 不存在", notExistIDs))
		return
	}

	err = global.DB.Debug().Delete(&logList).Error
	if err != nil {
		res.FailWithMsg("删除日志失败", c)
	}

	res.OkWithMsg(fmt.Sprintf("删除日志成功，共删除 %d 条", len(cr.IDs)), c)
}
