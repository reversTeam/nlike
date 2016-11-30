package main

import (
	"context"
	"flag"
	"github.com/reversTeam/nlike/server"
)

func main() {
	var host = flag.String("host", "127.0.0.1", "Default listening host")
	var port = flag.Int("port", 4242, "Default listening port")

	flag.Parse()

	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)

	server := nlike.NewServer(*host, *port)
	server.Start()
	defer server.Stop()
}
