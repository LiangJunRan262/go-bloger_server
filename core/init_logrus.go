package core

import (
	"bloger_server/global"
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Fromat 实现 Fromatter(entry *logrus.Entry) ([]byte, error) 接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的级别显示不同的颜色
	var levelColor int

	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, entry.Level, levelColor, entry.Level, entry.Message, fileVal, funcVal)
	} else {
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, entry.Level, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

type FileDateHook struct {
	file     *os.File
	logPath  string
	fileDate string // 文件名中的日期
	appName  string // 应用名称
}

func (hook FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook FileDateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02") // 日期格式
	line, _ := entry.String()                // 日志内容
	if hook.fileDate == timer {
		hook.file.Write([]byte(line)) // 写入文件
		return nil
	}

	// 时间不等
	hook.file.Close()                                                   // 关闭文件
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm) // 创建目录
	filename := fmt.Sprintf("%s/%s/%s.log", hook.logPath, timer, hook.appName)

	hook.file, _ = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 打开文件
	hook.fileDate = timer                                                           // 更新日期
	hook.file.Write([]byte(line))                                                   // 写入文件
	return nil
}

func InitFile(logPath, appName string) {
	fileDate := time.Now().Format("2006-01-02") // 日期格式
	// 创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm) // 创建目录
	if err != nil {
		logrus.Error(err) // 记录错误日志
		return
	}

	filename := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName) // 日志文件名

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) // 打开文件
	if err != nil {
		logrus.Error(err) // 记录错误日志
		return
	}

	fileHook := FileDateHook{file: file, logPath: logPath, fileDate: fileDate, appName: appName} // 创建钩子
	logrus.AddHook(&fileHook)                                                                    // 添加钩子
}

func InitLogrus() {
	logrus.SetOutput(os.Stdout)          // 标准输出
	logrus.SetReportCaller(true)         // 显示文件名和行号
	logrus.SetFormatter(&LogFormatter{}) // 自定义日志格式
	logrus.SetLevel(logrus.DebugLevel)   // 设置日志级别
	InitFile(global.Config.Log.Dir, global.Config.Log.App)
	return
}
