package validator

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"shop_api/order-web/global"
)

func InitCustomValidator() error {
	err := RegisterValidator()
	if err != nil {
		return err
	}
	return nil
}
func RegisterValidator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("mobile", ValidateMobile)
		if err != nil {
			return err
		}
		err = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
		if err != nil {
			return err
		}
	}
	return nil
}
