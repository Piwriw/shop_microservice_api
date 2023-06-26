package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"shop_api/good-web/global"
	"shop_api/good-web/initialize"
	"syscall"
)

func main() {
	router := initialize.Routers()
	if err := initialize.InitLogger(); err != nil {
		fmt.Printf("init Logger  failed, err:%v\n", err)
		return
	}
	if err := initialize.InitReadNacos(); err != nil {
		fmt.Printf("init InitReadNacos  failed, err:%v\n", err)
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
	registerClient, err, serverId := initialize.InitRegisterClient()
	if err != nil {
		fmt.Printf("init Consult RegisterClient  failed, err:%v\n", err)
	}
	zap.S().Infof("启动服务器，端口:%d", global.AppConf.System.Port)
	go func() {
		if err := router.Run(fmt.Sprintf("%s:%d", global.AppConf.System.Host, global.AppConf.System.Port)); err != nil {
			zap.S().Panic(err)
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = registerClient.DeRegister(serverId); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
