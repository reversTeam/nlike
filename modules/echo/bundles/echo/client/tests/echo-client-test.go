package main

import (
	"github.com/reversTeam/nlike/modules/echo/bundles/echo/client"
	"log"
)

func main() {
	nb_pistes := 400
	nb_request := 2500
	done := make(chan bool, nb_pistes*nb_request)

	c, err := client.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("START")
	for i := 0; i < nb_pistes; i++ {
		go func(c *client.Client, done chan bool) {
			for j := 0; j < nb_request; j++ {
				c.Echo("Hello world !")
				done <- true
			}
		}(c, done)
	}

	for j := 0; j < nb_pistes*nb_request; j++ {
		<-done
	}
	log.Println("END")

}
