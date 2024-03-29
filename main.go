package main

import (
	"flag"
	"github.com/reversTeam/nlike/kernel"
	"github.com/reversTeam/nlike/modules"
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
	defer log.Println("[MAIN]: Systeme is ready for catch exit's signals")

	go func() {
		sig := <-sigs
		defer log.Println("[SYSTEM]: Signal catch:", sig)
		done <- true
	}()
	return
}

func main() {
	defer log.Println("[MAIN]: End program")
	host, port := getFlags()
	done := configureSignals()

	srv := kernel.NewServer(*host, *port)
	modules.Init(srv)
	srv.InitPackages()
	srv.Start()
	defer srv.Stop()

	<-done
}
