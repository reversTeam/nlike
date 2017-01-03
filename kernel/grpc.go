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
	InitRequests()
	serve()
}

type Grpc struct {
	name     string
	host     string
	port     int
	Server   *grpc.Server
	listener net.Listener
	done     chan error
}

func NewGrpc(name string, host string, port int) (o *Grpc) {
	o = &Grpc{
		name: name,
		host: host,
		port: port,
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", o.host, o.port))
	if err != nil {
		log.Panic(err)
	}
	o.done = make(chan error)
	o.listener = lis
	o.Server = grpc.NewServer()

	return o
}

func (o *Grpc) Init() {
	o.Start()
}

func (o *Grpc) InitRequest() {
	log.Printf("[%s]: Init Request", o.name)
}

func (o *Grpc) Start() {
	defer log.Printf("[%s]: listen on %s:%d", o.name, o.host, o.port)

	go o.serve()
}

func (o *Grpc) Stop() {
	defer log.Printf("[%s]: Server stop", o.name)
	o.Server.GracefulStop()
}

func (o *Grpc) serve() {
	o.done <- o.Server.Serve(o.listener)
}
