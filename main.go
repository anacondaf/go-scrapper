package main

import (
	"log"

	"github.com/kainguyen/go-scrapper/src/core/application/http"
	"github.com/kainguyen/go-scrapper/src/infrastructure/di"
)

func main() {
	di.ContainerRegister()
	
	server, err := http.NewHttpServer()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = server.StartApp(":3000")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
