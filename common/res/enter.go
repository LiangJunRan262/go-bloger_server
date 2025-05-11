package res

import (
	"bloger_server/utils/validate"
	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	SuccessCode     ResponseCode = 0    // 成功
	FailValidCode   ResponseCode = 1001 // 校验失败
	ServerErrorCode ResponseCode = 1002 // 服务器错误
)

func (c ResponseCode) String() string {
	switch c {
	case SuccessCode:
		return "成功"
	case FailValidCode:
		return "参数校验失败"
	case ServerErrorCode:
		return "服务异常"
	default:
		return "unknown"
	}
}

var empty = map[string]any{}

type Response struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

func (r Response) Json(c *gin.Context) {
	c.JSON(200, r)
}

func Ok(data any, msg string, c *gin.Context) {
	Response{
		Code: SuccessCode,
		Msg:  msg,
		Data: data,
	}.Json(c)
}

func OkWithData(data any, c *gin.Context) {
	Response{
		Code: SuccessCode,
		Msg:  "成功",
		Data: data,
	}.Json(c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Response{
		Code: SuccessCode,
		Msg:  msg,
		Data: empty,
	}.Json(c)
}

func OkWithList(list any, count int, c *gin.Context) {
	Response{
		Code: SuccessCode,
		Msg:  "成功",
		Data: map[string]any{
			"list":  list,
			"count": count,
		},
	}.Json(c)
}

// 失败
func FailCode(code ResponseCode, msg string, c *gin.Context) {
	Response{
		Code: code,
		Msg:  msg,
		Data: empty,
	}.Json(c)
}

// 参数校验失败
func FailValid(msg string, data any, c *gin.Context) {
	Response{
		Code: FailValidCode,
		Msg:  msg,
		Data: data,
	}.Json(c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Response{
		Code: ServerErrorCode,
		Msg:  msg,
		Data: empty,
	}.Json(c)
}

func FailWithData(data any, msg string, c *gin.Context) {
	Response{
		Code: ServerErrorCode,
		Msg:  "失败哦",
		Data: data,
	}.Json(c)
}

func FailWidthError(err error, c *gin.Context) {
	data, msg := validate.ValidateError(err)
	FailValid(msg, data, c)
}
