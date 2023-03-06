package main

import (
	"github.com/kainguyen/go-scrapper/src/core/application/http"
	"log"
)

func main() {
	server, err := http.NewHttpServer()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = server.StartApp(":3000")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
