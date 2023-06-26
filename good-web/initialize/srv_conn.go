package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shop_api/good-web/global"
	"shop_api/good-web/proto"
)

// InitSrvConn 配置了负载均衡的grpc client
func InitSrvConn() error {
	consulInfo := global.AppConf.Consul
	userConn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.IP, consulInfo.Port, global.AppConf.GrpcServer.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 【商品服务失败】", "msg", err.Error())
		return err
	}
	global.GoodSrvClient = proto.NewGoodsClient(userConn)
	return nil
}
