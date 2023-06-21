package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	R "shop_api/user-web/api/response"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		R.ResponseError(c, R.CodeCreateVerificationCode)
		return
	}
	R.ResponseSuccess(c, &R.Captchcha{
		CaptchchaID: id,
		PicPath:     b64s,
	})
}
