package grpc

import (
	"google.golang.org/grpc"
	"log"
)

type RequestInterface interface {
	InitServices(s *grpc.Server)
}

type Request struct {
}

func NewRequest() {

}

func (o *Request) InitServices(s *grpc.Server) {
	log.Println("[REQUEST] Init Services")
}
