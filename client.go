package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	microdemo_service_gomicro "gomicro-demo/proto/gomicro"
)

/**
go-micro客户端程序
启动命令(命令行中指定etcd服务的地址)：go run ./client.go --registry=etcd --registry_address=localhost:2379 --server_address=:9100
命令行调用micro服务服务的方法：micro call microdemo.service.gomicro Gomicro.Call '{"name": "John"}
*/
func main() {
	// create the greeter client using the service name and client
	service := micro.NewService(
		micro.Name("microdemo.service.gomicro"),
	//micro.Version("v0.1.0"),
	)
	service.Init()
	//创建客户端
	gomicro := microdemo_service_gomicro.NewGomicroService("microdemo.service.gomicro", service.Client())

	//调用rpc服务
	rsp, err := gomicro.Call(context.TODO(), &microdemo_service_gomicro.Request{Name: "helloworld"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.GetMsg())
}
