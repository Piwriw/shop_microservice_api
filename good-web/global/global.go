package global

import (
	ut "github.com/go-playground/universal-translator"
	"shop_api/good-web/proto"
)

var (
	AppConf       AppConfig
	NacosConf     Nacos
	Trans         ut.Translator
	GoodSrvClient proto.GoodsClient
)
