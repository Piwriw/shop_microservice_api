package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shop_api/order-web/global"
	"shop_api/order-web/proto"
)

func getSrvConn(serverName string) (*grpc.ClientConn, error) {
	consulInfo := global.AppConf.Consul
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.IP, consulInfo.Port, serverName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 【订单服务失败】", "msg", err.Error())
		return nil, err
	}
	return conn, nil
}

// InitSrvConn 配置了负载均衡的grpc client
func InitSrvConn() error {
	orderConn, err := getSrvConn(global.AppConf.OrderGrpcServer.Name)
	if err != nil {
		return err
	}
	global.OrderSrvClient = proto.NewOrderClient(orderConn)

	goodConn, err := getSrvConn(global.AppConf.GoodGrpcServer.Name)
	if err != nil {
		return err
	}
	global.GoodSrvClient = proto.NewGoodsClient(goodConn)

	invConn, err := getSrvConn(global.AppConf.InventoryGrpcServer.Name)
	if err != nil {
		return err
	}
	global.InventoryClient = proto.NewInventoryClient(invConn)
	return nil
}
