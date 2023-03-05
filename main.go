package main

import (
	"github.com/kainguyen/go-scrapper/src/core/application/api"
	"github.com/kainguyen/go-scrapper/src/infrastructure/webScraping"
	"log"
)

func main() {
	var goColly = webScraping.NewColly()

	server, err := api.NewHttpServer(goColly)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = server.StartApp(":3000")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
