package grpc

import (
	"github.com/reversTeam/nlike/kernel"
	echoRequest "github.com/reversTeam/nlike/modules/echo/bundles/echo/grpc/request"
)

type EchoGrpc struct {
	kernel.Grpc
}

func NewGrpc() *EchoGrpc {
	grpc := &EchoGrpc{
		*kernel.NewGrpc("EchoGrpc", "127.0.0.1", 4244),
	}

	return grpc
}

func (o *EchoGrpc) InitRequests() {
	echoReq := echoRequest.NewRequest()
	echoReq.InitServices(o.Server)
}
