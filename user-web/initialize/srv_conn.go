package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shop_api/user-web/global"
	"shop_api/user-web/proto"
)

// InitSrvConn 配置了负载均衡的grpc client
func InitSrvConn() error {
	consulInfo := global.AppConf.Consul
	userConn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.IP, consulInfo.Port, global.AppConf.System.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 【用户服务失败】", "msg", err.Error())
		return err
	}
	global.UserSrvClient = proto.NewUserClient(userConn)
	return nil
}

func InitSrvConn2() error {
	// 从注册服务中心拉取
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.AppConf.Consul.IP, global.AppConf.Consul.Port)
	userSrcHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Fatalf("[InitSrvConn] 查询 【用户列表失败,拉取服务失败】err:%s", err.Error())
		return err
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service==%s`, global.AppConf.System.Name))
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 查询 【查询User服务失败】err:%s", err.Error())
		return err
	}
	for _, service := range data {
		userSrcHost = service.Address
		userSrvPort = service.Port
	}
	if userSrcHost == "" {
		return err
	}

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrcHost, userSrvPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", setting.Conf.GrpcServer.Host, setting.Conf.GrpcServer.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 【用户服务失败】", "msg", err.Error())
		return err
	}
	global.UserSrvClient = proto.NewUserClient(userConn)
	return nil
}
