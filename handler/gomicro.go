package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	gomicro "gomicro-demo/proto/gomicro"
)

type Gomicro struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Gomicro) Call(ctx context.Context, req *gomicro.Request, rsp *gomicro.Response) error {
	log.Info("Received Gomicro.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Gomicro) Stream(ctx context.Context, req *gomicro.StreamingRequest, stream gomicro.Gomicro_StreamStream) error {
	log.Infof("Received Gomicro.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&gomicro.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Gomicro) PingPong(ctx context.Context, stream gomicro.Gomicro_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&gomicro.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
