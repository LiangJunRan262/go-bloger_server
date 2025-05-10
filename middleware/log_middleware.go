package middleware

import (
	log_service "bloger_server/service/log_servive"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
	Head http.Header
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	w.Body = append(w.Body, b...)
	return w.ResponseWriter.Write(b)
}

func (w *ResponseWriter) Header() http.Header {
	return w.Head
}

func LogMiddleware(c *gin.Context) {
	log := log_service.NewActionLogByGin(c)

	// 在请求处理之前执行的逻辑
	// 可以在这里记录日志、验证权限等
	log.SetRequest(c)

	c.Set("log", log)

	res := &ResponseWriter{
		ResponseWriter: c.Writer,
		Body:           []byte{},
		Head:           make(http.Header),
	}

	c.Writer = res
	c.Next() // 调用下一个中间件或处理函数

	// 在请求处理之后执行的逻辑
	// 可以在这里记录响应时间、处理错误等
	log.SetResponse(res.Body)
	log.SetResponseHeader(res.Head)

	// 保存日志到数据库
	// fmt.Println("日志保存成功", string(log))

	log.Save()
}
