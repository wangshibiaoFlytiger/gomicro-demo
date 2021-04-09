package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"gomicro-demo/handler"
	"gomicro-demo/subscriber"

	gomicro "gomicro-demo/proto/gomicro"
)

/**
go-micro服务
启动方法：go run ./main.go  --registry=etcd --registry_address=localhost:2379
*/
func main() {
	// New Service
	service := micro.NewService(
		micro.Name("microdemo.service.gomicro"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	gomicro.RegisterGomicroHandler(service.Server(), new(handler.Gomicro))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("microdemo.service.gomicro", service.Server(), new(subscriber.Gomicro))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
