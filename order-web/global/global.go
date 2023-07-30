package global

import (
	ut "github.com/go-playground/universal-translator"
	"shop_api/order-web/proto"
)

var (
	AppConf         AppConfig
	NacosConf       Nacos
	Trans           ut.Translator
	OrderSrvClient  proto.OrderClient
	GoodSrvClient   proto.GoodsClient
	InventoryClient proto.InventoryClient
)
