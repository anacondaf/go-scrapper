package main

import (
	"github.com/kainguyen/go-scrapper/src/core/application/http"
	DIContainer "github.com/kainguyen/go-scrapper/src/infrastructure/di"
	"log"
)

func main() {
	DIContainer.ContainerRegister()

	server, err := http.NewHttpServer()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = server.StartApp(":3000")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
