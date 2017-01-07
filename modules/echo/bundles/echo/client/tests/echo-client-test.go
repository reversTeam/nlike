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

func getFlags() (log_level *int, nb_concurents *int, nb_requests *int) {
	nb_concurents = flag.Int("concurents", 400, "Number of concourrents")
	nb_requests = flag.Int("requests", 5000, "Number of requests")
	log_level = flag.Int("log", 1, "Display log mode")

	flag.Parse()
	return
}

func main() {
	log_level, nb_concurents, nb_requests := getFlags()
	total_requests := *nb_concurents * *nb_requests

	log.Printf("[%s]: Start benchmark", name)
	defer log.Printf("[%s]: End benchmark", name)
	log.Printf("[%s]: intiialized with %d level logs", name, *log_level)
	log.Printf("[%s]: initialized with %d concurents", name, *nb_concurents)
	log.Printf("[%s]: initialized with %d requets", name, *nb_requests)
	wg := sync.WaitGroup{}

	c, err := client.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	wg.Add(*nb_concurents)
	for i := 0; i < *nb_concurents; i++ {
		go func(c *client.Client, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := 0; j < *nb_requests; j++ {
				c.Echo("Hello world !")
			}
		}(c, &wg)
	}
	wg.Wait()
	past_time := time.Since(start)
	past_time_in_nano := float64(past_time.Nanoseconds())
	past_time_in_seconds := past_time_in_nano / 1000000000
	requests_per_second := float64(total_requests) / past_time_in_seconds

	avg_time_per_request := past_time_in_nano / float64(total_requests)

	log.Printf("[%s]: proceded %d in %s", name, total_requests, past_time)
	log.Printf("[%s]: result %.3f requests per seconds", name, requests_per_second)
	log.Printf("[%s]: result %.0f avg nanoseconds per request", name, avg_time_per_request)

}
