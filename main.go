package main

import (
	"fmt"
	"go.uber.org/zap"
	"shop_api/user-web/global"
	"shop_api/user-web/initialize"
	myvalidator "shop_api/user-web/validator"
)

func main() {
	router := initialize.Routers()
	if err := initialize.InitLogger(); err != nil {
		fmt.Printf("init Logger  failed, err:%v\n", err)
		return
	}
	if err := initialize.InitReadNacos(); err != nil {
		fmt.Printf("init InitSetting  failed, err:%v\n", err)
		return
	}
	if err := initialize.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	if err := initialize.InitSrvConn(); err != nil {
		fmt.Printf("init srvconn  failed, err:%v\n", err)
		return
	}
	if err := myvalidator.InitCustomValidator(); err != nil {
		fmt.Printf("validator InitCustomValidator failed, err:%v\n", err)
		return
	}
	zap.S().Info("启动服务器，端口:%d", global.AppConf.System.Port)
	if err := router.Run(fmt.Sprintf("%s:%d", global.AppConf.System.Host, global.AppConf.System.Port)); err != nil {
		zap.S().Panic(err)
	}
}
