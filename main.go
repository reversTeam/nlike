package main

import (
	"flag"
	"github.com/reversTeam/nlike/server"
	"github.com/reversTeam/nlike/topic"
	"log"
	"net/http"

	"os"
	"os/signal"
	"syscall"
)

func getFlags() (host *string, port *int) {
	host = flag.String("host", "127.0.0.1", "Default listening host")
	port = flag.Int("port", 4242, "Default listening port")
	flag.Parse()

	return
}

func configureSignals() (done chan bool) {
	done = make(chan bool, 1)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer log.Println("Systeme is ready for catch exit's signals")

	go func() {
		sig := <-sigs
		defer log.Println("Signal catch:", sig)
		done <- true
	}()
	return
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

// func addTopicRoute(server *nlike.Server) {
// 	server.AddRoute("/topics", handler)
// 	server.AddRoute("/topics/create", handler)
// 	server.AddRoute("topics/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$", handler)
// 	server.AddRoute("topics/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}/update$", handler)
// 	server.AddRoute("topics/[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}/delete$", handler)
// }

func main() {
	defer log.Println("End program")
	host, port := getFlags()
	done := configureSignals()

	srv := server.NewServer(*host, *port)
	topic.NewTopic()

	srv.Start()
	defer srv.Stop()

	<-done
}
