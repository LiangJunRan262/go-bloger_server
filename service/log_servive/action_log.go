package log_service

import (
	"bloger_server/core"
	"bloger_server/global"
	"bloger_server/models"
	"bloger_server/models/enum"
	"bloger_server/utils/jwts"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	e "github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ItemType struct {
	Level string `json:"level"`
	Label string `json:"label"`
	Value string `json:"value"`
}

// 定义请求结构体
type ContentType struct {
	Method         string     `json:"method"`
	URL            string     `json:"url"`
	ItemList       []ItemType `json:"itemList"`
	RequestBody    string     `json:"requestBody"`
	RequestHeader  string     `json:"requestHeader"`
	ResponseBody   string     `json:"responseBody"`
	ResponseHeader string     `json:"responseHeader"`
	ErrMsg         string     `json:"errMsg"`
}

type ActionLog struct {
	c                  *gin.Context
	level              enum.LogLevelType
	title              string
	requestBody        []byte
	responseBody       []byte
	log                *models.LogModel
	showRequest        bool
	showRequestHeader  bool
	showResponse       bool
	showResponseHeader bool
	content            ContentType
	isMiddlewareSave   bool
}

func (c *ActionLog) SetShowRequest() {
	c.showRequest = true
}

func (c *ActionLog) SetShowRequestHeader() {
	c.showRequestHeader = true
}

func (c *ActionLog) SetShowResponse() {
	c.showResponse = true
}

func (c *ActionLog) SetShowResponseHeader() {
	c.showResponseHeader = true
}

func (ac *ActionLog) SetLevel(level enum.LogLevelType) {
	ac.level = level
}

func (ac *ActionLog) SetTitle(title string) {
	ac.title = title
}

func (ac *ActionLog) SetItem(label string, value any, LogLevelType enum.LogLevelType) {
	v := ""
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		byteData, _ := json.Marshal(value)
		v = string(byteData)
	default:
		v = fmt.Sprintf("%v", value)
	}
	ac.content.ItemList = append(ac.content.ItemList, ItemType{
		Level: LogLevelType.String(),
		Label: label,
		Value: v,
	})
}
func (ac *ActionLog) SetItemInfo(label string, value any) {
	ac.SetItem(label, value, enum.LogInfoLevel)
}
func (ac *ActionLog) SetItemWarn(label string, value any) {
	ac.SetItem(label, value, enum.LogWarnLevel)
}
func (ac *ActionLog) SetItemErr(label string, value any) {
	// 如果value是error类型
	if err, ok := value.(error); ok {
		msg := e.WithStack(err)
		stackTrace := fmt.Sprintf("%+v\n", msg)
		value = stackTrace
	}
	logrus.Errorf("参数绑定错误: %+v\n", value)
	ac.SetItem(label, value, enum.LogErrLevel)
}

// 设置请求体
func (ac *ActionLog) SetRequest(c *gin.Context) {
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorf("读取请求体错误: %v", err.Error())
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteData))
	ac.requestBody = byteData
}

// 设置响应体
func (ac *ActionLog) SetResponse(data []byte) {
	ac.responseBody = data
}

func (ac *ActionLog) SetResponseHeader(header http.Header) {
	json, _ := json.Marshal(header)
	ac.content.ResponseHeader = string(json)
}

// 专门给gin使用的中间件
func (ac *ActionLog) MiddlewareSave() {
	if ac.log == nil {
		ac.isMiddlewareSave = true
		ac.Save()
		return
	}
	// 在视图里save过，属于更新
	// 设置响应
	if ac.showResponse {
		ac.content.ResponseBody = string(ac.responseBody)
	}
}

// TODO: 当前log.Sava() 只在中间件最后执行， 之后优化
func (ac *ActionLog) Save() (id uint) {
	if ac.log == nil && ac.c.GetBool("saveLog") == false {
		return 0
	}

	// TODO: 这里需要判断是否是中间件保存的，因为中间件保存的话，需要在视图里保存
	// 若已存在日志，则更新日志内容，否则创建新日志
	if ac.log != nil {
		return ac.log.ID
	}

	// 设置请求
	if ac.showRequest {
		ac.content.RequestBody = string(ac.requestBody)
	}
	if ac.showRequestHeader {
		json, _ := json.Marshal(ac.c.Request.Header)
		ac.content.RequestHeader = string(json)
	}
	if ac.isMiddlewareSave {
		// 设置响应
		if ac.showResponse {
			ac.content.ResponseBody = string(ac.responseBody)
		}
	}

	contentStr, contentStrErr := json.Marshal(ac.content)
	if contentStrErr != nil {
		logrus.Errorf("序列化操作日志失败 %s", contentStrErr.Error())
	}

	ip := ac.c.ClientIP()
	addr := core.GetIPAddr(ip)
	claims, err := jwts.ParseTokenByGin(ac.c)
	userID := uint(0)
	if err == nil && claims != nil {
		userID = claims.UserID
	}

	log := models.LogModel{
		LogType: enum.ActionLogType,
		Title:   ac.title,
		Content: string(contentStr),
		Level:   ac.level,
		UserID:  userID, // 假设从token中获取用户ID，这里简化为1
		IP:      ip,
		Address: addr,
	}

	err = global.DB.Create(&log).Error
	if err != nil {
		logrus.Errorf("保存操作日志失败 %s", err.Error())
	}
	ac.log = &log

	return log.ID
}

func NewActionLogByGin(c *gin.Context) *ActionLog {
	return &ActionLog{
		c: c,
		content: ContentType{
			Method:         c.Request.Method,
			URL:            c.Request.URL.String(),
			ItemList:       []ItemType{},
			RequestBody:    "",
			ResponseHeader: "",
			ResponseBody:   "",
			RequestHeader:  "",
		},
	}
}

func GetActionLog(c *gin.Context) *ActionLog {
	c.Set("saveLog", true)
	ac, exists := c.Get("log")
	if !exists {
		return NewActionLogByGin(c)
	}
	return ac.(*ActionLog)
}
