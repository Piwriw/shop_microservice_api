package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	R "shop_api/user-web/api/response"
	"shop_api/user-web/global"
	"strings"
)

// removeTopStruct 去除提示信息中的结构体名称
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		R.FailWithDetailed(c, "参数验证不通过", err.Error())
		return
	}
	R.FailWithDetailed(c, "参数验证不通过", removeTopStruct(errs.Translate(global.Trans)))
	return
}
