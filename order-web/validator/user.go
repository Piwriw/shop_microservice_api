package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(?:(?:\+|00)86)?1[3456789]\d{9}$`, mobile)
	if !ok {
		return false
	}
	return true
}
