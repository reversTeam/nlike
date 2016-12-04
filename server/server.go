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

func (o *Server) Start() {
	tmp := fmt.Sprintf("%s:%d", o.host, o.port)

	log.Println("Server started and listen:", tmp)
	go func() {
		defer o.Stop()
		if err := http.ListenAndServe(tmp, nil); err != nil {
			log.Panic("Server ERROR[500]:", err)
		}
	}()
}

func (o *Server) Stop() {
	log.Println("Server stop")
}
