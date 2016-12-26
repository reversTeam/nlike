package grpc

import (
	nlike "github.com/reversTeam/nlike/kernel"
)

type TopicGrpc struct {
	nlike.Grpc
}

func NewGrpc() *TopicGrpc {
	grpc := &TopicGrpc{
		*nlike.NewGrpc("TopicGrpc", "127.0.0.1", 4243),
	}

	return grpc
}

func (o *TopicGrpc) InitRequests() {

}
