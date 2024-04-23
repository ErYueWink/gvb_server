package res

import "github.com/gin-gonic/gin"

const (
	SUCCESS = 0
	ERROR   = 7
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type ListResponse[T any] struct {
	List  []T   `json:"list"`
	Count int64 `json:"count"`
}

// Result 请求结果
func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(code, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// OK 请求成功 返回数据
func OK(data any, msg string, c *gin.Context) {
	Result(SUCCESS, msg, data, c)
}

// OKWithMsg 请求成功有消息返回
func OKWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, msg, map[string]interface{}{}, c)
}

// OKWithData 请求成功有数据返回
func OKWithData(data any, c *gin.Context) {
	Result(SUCCESS, "请求成功", data, c)
}

// OKWith 请求成功
func OKWith(c *gin.Context) {
	Result(SUCCESS, "请求成功", map[string]interface{}{}, c)
}

// Fail 请求失败
func Fail(code int, msg string, c *gin.Context) {
	Result(code, msg, map[string]interface{}{}, c)
}

// FailWithMsg 请求失败有消息返回
func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, msg, map[string]interface{}{}, c)
}

// FailErrorCode 请求失败后返回错误码 业务：参数绑定失败等等
func FailErrorCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), msg, map[string]interface{}{}, c)
		return
	}
	Result(int(code), "系统未知错误，请反馈给管理员", map[string]interface{}{}, c)
}
