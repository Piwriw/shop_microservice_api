package global

import (
	ut "github.com/go-playground/universal-translator"
	"shop_api/user-web/proto"
)

var (
	AppConf       AppConfig
	NacosConf     Nacos
	Trans         ut.Translator
	UserSrvClient proto.UserClient
)
