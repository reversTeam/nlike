package nlike

import (
	"fmt"
	"log"
	"net/http"
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
	tmp := fmt.Sprintf("%s:%d", o.host, o.port)

	log.Println("Server started and listen:", tmp)
	return http.ListenAndServe(tmp, nil)
}

func (o *Server) Stop() {
	log.Println("Server stop")
}
