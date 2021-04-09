package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	gomicro "gomicro-demo/proto/gomicro"
)

type Gomicro struct{}

func (e *Gomicro) Handle(ctx context.Context, msg *gomicro.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *gomicro.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
