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

func RespError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}
func RespErrorWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 500,
		Msg:  msg,
		Data: nil,
	})
}
func RespSuccessMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 200,
		Msg:  msg,
		Data: nil,
	})
}
func RespSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func FailWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  msg,
		"data": data,
	})
}
