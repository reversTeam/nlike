package grpc

import (
	"github.com/reversTeam/nlike/kernel"
)

type TopicGrpc struct {
	kernel.Grpc
}

func NewGrpc() *TopicGrpc {
	grpc := &TopicGrpc{
		*kernel.NewGrpc("127.0.0.1", 4243),
	}

	return grpc
}
