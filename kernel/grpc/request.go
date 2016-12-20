package grpc

import (
	ggrpc "google.golang.org/grpc"
	"log"
)

type RequestInterface interface {
	InitServices(s *ggrpc.Server)
}

type Request struct {
}

func NewRequest() {

}

func (o *Request) InitServices(s *ggrpc.Server) {
	log.Println("[REQUEST] Init Services")
}
