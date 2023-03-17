package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/kainguyen/go-scrapper/src/core/application/http"
	_ "github.com/kainguyen/go-scrapper/src/infrastructure/di"
	"github.com/kainguyen/go-scrapper/src/utils"
	"go/parser"
	"go/token"
	"log"
)

func init() {
	fs := token.NewFileSet()

	fileName := utils.GetWorkDirectory()

	f, err := parser.ParseFile(fs, fmt.Sprintf("%v/src/core/application/http/post/post.handler.go", fileName), nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	//
	//fmt.Printf("%#v\n", f)

	spew.Dump(f)
}

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
