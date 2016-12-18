package grpc

import (
	"github.com/reversTeam/nlike/kernel"
)

type EchoGrpc struct {
	kernel.Grpc
}

func NewGrpc() *EchoGrpc {
	grpc := &EchoGrpc{
		*kernel.NewGrpc("127.0.0.1", 4244),
	}

	return grpc
}

func (o *EchoGrpc) InitProtos() {

}
