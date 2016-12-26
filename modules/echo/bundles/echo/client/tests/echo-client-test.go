package main

import (
	"github.com/reversTeam/nlike/modules/echo/bundles/echo/client"
	"log"
)

func main() {
	c, err := client.NewClient()

	if err != nil {
		log.Fatal(err)
	}
	c.Echo("Hello world !")
}
