package nlike

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

type Server struct {
	host string
	port int
}

func NewServer(host string, port int) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

func (o *Server) Start() error {
	tmp := fmt.Sprintf("%s:%s", o.host, strconv.Itoa(o.port))

	log.Println("Server started and listen:", tmp)
	_, err := net.Listen("tcp", tmp)
	if err != nil {
		return err
	}
	return nil
}

func (o *Server) Stop() {
	log.Println("Server stop")
}
