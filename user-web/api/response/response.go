package R

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	success = 200
	fail    = 500
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  "成功！",
	})
}
func SuccessWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  msg,
	})
}
func SuccessWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  msg,
		"data": data,
	})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  "失败！",
	})
}
func FailWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  msg,
	})
}
func FailWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  msg,
		"data": data,
	})
}
