package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func FilterService() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.28.145:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(`Service=="user-web"`)
	for k, _ := range data {
		fmt.Println(k)
	}
}
func AllService() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.28.145:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().Services()
	for k, _ := range data {
		fmt.Println(k)
	}
}
func Register(address, name string, port int, tags []string, id string) {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.28.145:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://127.0.0.1:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10",
	}
	// 生成注册对象
	registerion := &api.AgentServiceRegistration{
		Name:    name,
		ID:      id,
		Port:    port,
		Tags:    tags,
		Address: address,
		Check:   check,
	}
	err = client.Agent().ServiceRegister(registerion)
	if err != nil {
		panic(err)
	}
}
