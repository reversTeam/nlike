package main

import (
	"fmt"
	"github.com/reversTeam/nlike/modules/echo/bundles/echo/client"
	"log"
	"sync"
	"time"
)

func main() {
	nb_pistes := 400
	nb_request := 2500
	wg := sync.WaitGroup{}

	c, err := client.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	wg.Add(nb_pistes)
	for i := 0; i < nb_pistes; i++ {
		go func(c *client.Client, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := 0; j < nb_request; j++ {
				c.Echo("Hello world !")
			}
		}(c, &wg)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
