package kernel

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcInterface interface {
	Init()
	Start()
	Stop()
	InitProtos()
	serve()
}

type Grpc struct {
	host     string
	port     int
	server   *grpc.Server
	listener net.Listener
	done     chan error
}

func NewGrpc(host string, port int) *Grpc {
	return &Grpc{
		host: host,
		port: port,
	}
}

func (o *Grpc) Init() {
	o.Start()
}

func (o *Grpc) InitProtos() {

}

func (o *Grpc) Start() {
	defer log.Printf("Server GRPC listen on %s:%d", o.host, o.port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", o.host, o.port))
	if err != nil {
		log.Panic(err)
	}
	o.server = grpc.NewServer()
	o.listener = lis
	o.done = make(chan error)
	go o.serve()
}

func (o *Grpc) Stop() {
	defer log.Println("Server GRPC stop")
	o.server.GracefulStop()
}

func (o *Grpc) serve() {
	o.done <- o.server.Serve(o.listener)
}
