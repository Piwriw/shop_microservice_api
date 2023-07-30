package initialize

import (
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"shop_api/order-web/global"
	"shop_api/order-web/utils/register/consul"
)

func InitRegisterClient() (consul.RegistryClient, error, string) {
	registerClient := consul.NewRegistryClient(global.AppConf.Consul.IP, global.AppConf.Consul.Port)
	u, err := uuid.NewV4()
	if err != nil {
		panic("failed new  uuid:" + err.Error())
	}
	serverId := u.String()
	err = registerClient.Register(global.AppConf.System.Host, global.AppConf.System.Port,
		global.AppConf.System.Name, global.AppConf.System.Tags, serverId)
	if err != nil {
		return nil, err, ""
	}
	return registerClient, nil, serverId
}
