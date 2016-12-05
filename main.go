package main

import (
	"flag"
	"github.com/reversTeam/nlike/server"
	"github.com/reversTeam/nlike/topic"
	"log"

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

func main() {
	defer log.Println("End program")
	host, port := getFlags()
	done := configureSignals()

	srv := server.NewServer(*host, *port)
	tpc := topic.NewTopic()
	tpc.Init(srv.GetRouter())

	srv.Start()
	defer srv.Stop()

	<-done
}
