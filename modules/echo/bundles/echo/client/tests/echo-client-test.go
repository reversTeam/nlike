package main

import (
	"flag"
	"github.com/reversTeam/nlike/modules/echo/bundles/echo/client"
	"log"
	"sync"
	"time"
)

const (
	name = "EchoClientBenchmark"
)

func getFlags() (
	log_level *int,
	host *string,
	port *int,
	nb_concurents *int,
	nb_requests *int,
	nb_clients *int,
	message *string,
) {
	host = flag.String("host", "127.0.0.1", "Host")
	port = flag.Int("port", 4244, "Port")
	nb_concurents = flag.Int("concurents", 400, "Number of concourrents")
	nb_clients = flag.Int("clients", 4, "Number of client connexion")
	nb_requests = flag.Int("requests", 5000, "Number of requests")
	log_level = flag.Int("log", 1, "Display log mode")
	message = flag.String("message", "Hello world", "Message send")

	flag.Parse()
	return
}

func main() {
	log_level, host, port, nb_concurents, nb_requests, nb_clients, message := getFlags()
	total_requests := *nb_concurents * *nb_requests

	log.Printf("[%s]: Start benchmark", name)
	defer log.Printf("[%s]: End benchmark", name)
	log.Printf("[%s]: intiialized with \"%d\" level logs", name, *log_level)
	log.Printf("[%s]: intiialized with \"%s\" host", name, *host)
	log.Printf("[%s]: intiialized with \"%d\" port", name, *port)
	log.Printf("[%s]: initialized with \"%d\" clients connexion", name, *nb_clients)
	log.Printf("[%s]: initialized with \"%d\" concurents", name, *nb_concurents)
	log.Printf("[%s]: initialized with \"%d\" requets", name, *nb_requests)
	log.Printf("[%s]: initialized with \"%s\" message", name, *message)
	log.Printf("[%s]: start benchmark with %d requets", name, total_requests)

	wg := sync.WaitGroup{}

	clients := make([]*client.Client, 0)
	for i := 0; i < *nb_clients; i++ {
		c, err := client.NewClient(*host, *port)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, c)
	}

	start := time.Now()
	wg.Add(*nb_concurents)
	for i := 0; i < *nb_concurents; i++ {
		go func(c *client.Client, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := 0; j < *nb_requests; j++ {
				c.Echo(*message)
			}
		}(clients[i%*nb_clients], &wg)
	}
	wg.Wait()
	past_time := time.Since(start)
	past_time_in_nano := float64(past_time.Nanoseconds())
	avg_time_per_request := past_time_in_nano / float64(total_requests)
	requests_per_second := 1000000000 / avg_time_per_request

	log.Printf("[%s]: proceded %d in %s", name, total_requests, past_time)
	log.Printf("[%s]: result %.0f requests per seconds", name, requests_per_second)
	log.Printf("[%s]: result %.0f avg nanoseconds per request", name, avg_time_per_request)

}
