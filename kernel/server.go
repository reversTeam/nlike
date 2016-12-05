package kernel

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	host   string
	port   int
	router *Router
}

type config struct {
	filepath string
}

func NewServer(host string, port int) *Server {
	return &Server{
		host: host,
		port: port,
		// packages: make([]*Package, 0),
		router: NewRouter(),
	}
}

func (o *Server) Start() {
	tmp := fmt.Sprintf("%s:%d", o.host, o.port)

	defer log.Println("Server started and listen:", tmp)
	go func() {
		defer o.Stop()
		if err := http.ListenAndServe(tmp, o.router); err != nil {
			log.Panic("Server ERROR[500]:", err)
		}
	}()
}

func (o *Server) Stop() {
	defer log.Println("Server stop")
	// Execute action before closing the server
}

func (o *Server) GetRouter() *Router {
	return o.router
}
