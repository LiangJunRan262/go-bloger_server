package log_service

import (
	"bloger_server/global"
	"bloger_server/models"
	"bloger_server/models/enum"
	"encoding/json"
	"fmt"
	e "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"time"
)

type RuntimeLogItemType struct {
	Level string    `json:"level"`
	Label string    `json:"label"`
	Value string    `json:"value"`
	Time  time.Time `json:"time"`
}

type RuntimeLog struct {
	service         string               // 服务器名
	level           enum.LogLevelType    // 日志类型 1:登录 2:操作 3
	title           string               // 标题
	content         []RuntimeLogItemType // 内容
	runtimeDateType RunTimeDateType
}

func (ac *RuntimeLog) SetLevel(level enum.LogLevelType) {
	ac.level = level
}

func (ac *RuntimeLog) SetTitle(title string) {
	ac.title = title
}

func (ac *RuntimeLog) SetItem(label string, value any, LogLevelType enum.LogLevelType) {
	v := ""
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		byteData, _ := json.Marshal(value)
		v = string(byteData)
	default:
		v = fmt.Sprintf("%v", value)
	}
	ac.content = append(ac.content, RuntimeLogItemType{
		Level: LogLevelType.String(),
		Label: label,
		Value: v,
		Time:  time.Now(),
	})
}
func (ac *RuntimeLog) SetItemInfo(label string, value any) {
	ac.SetItem(label, value, enum.LogInfoLevel)
}
func (ac *RuntimeLog) SetItemWarn(label string, value any) {
	ac.SetItem(label, value, enum.LogWarnLevel)
}
func (ac *RuntimeLog) SetItemErr(label string, value any) {
	// 如果value是error类型
	if err, ok := value.(error); ok {
		msg := e.WithStack(err)
		stackTrace := fmt.Sprintf("%+v\n", msg)
		value = stackTrace
	}
	logrus.Errorf("参数绑定错误: %+v\n", value)
	ac.SetItem(label, value, enum.LogErrLevel)
}

// save方法
func (ac *RuntimeLog) Save() {
	// 判断是创建还是更新
	var log models.LogModel

	global.DB.Find(&log, fmt.Sprintf("service = ? and log_type = ? and created_at >= date_sub(now(), %s)", ac.runtimeDateType.GetSqlTime()), ac.service, enum.RuntimeLogType)

	contentStr, contentStrErr := json.Marshal(ac.content)
	if contentStrErr != nil {
		logrus.Errorf("序列化运行日志失败 %s", contentStrErr.Error())
	}

	if log.ID != 0 {
		// TODO 更新
		return
	}

	err := global.DB.Create(&models.LogModel{
		LogType: enum.RuntimeLogType,
		Title:   ac.title,
		Content: string(contentStr),
		Level:   ac.level,
		Service: ac.service,
	}).Error
	if err != nil {
		logrus.Errorf("保存运行日志失败 %s", err.Error())
	}

}

type RunTimeDateType int8

const (
	RunTimeDateHour  RunTimeDateType = iota + 1 // 小时
	RunTimeDateDay                              // 天
	RunTimeDateWeek                             // 周
	RunTimeDateMonth                            // 月
	RunTimeDateYear                             // 年
)

func (r RunTimeDateType) GetSqlTime() string {
	switch r {
	case RunTimeDateHour:
		return "interval 1 hour"
	case RunTimeDateDay:
		return "interval 1 day"
	case RunTimeDateWeek:
		return "interval 1 week"
	case RunTimeDateMonth:
		return "interval 1 month"
	case RunTimeDateYear:
		return "interval 1 year"
	default:
		return "interval 1 hour"
	}
}

func NewRuntimeLog(service string, dateType RunTimeDateType) *RuntimeLog {
	return &RuntimeLog{
		service:         service,
		runtimeDateType: dateType,
	}
}
