package grpc

import (
	ggrpc "google.golang.org/grpc"
	"log"
)

type RequestInterface interface {
	InitServices(s *ggrpc.Server)
}

type Request struct {
	name string
}

func NewRequest(name string) *Request {
	return &Request{
		name: name,
	}
}

func (o *Request) InitServices(s *ggrpc.Server) {
	log.Printf("[%s] Init Services", o.name)
}
