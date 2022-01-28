package main

import (
	"log"

	simpleauth "github.com/jast-r/simple-auth"
	"github.com/jast-r/simple-auth/pkg/handler"
)

func main() {
	server := new(simpleauth.Server)
	handlers := new(handler.Handler)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("run server failed: %s", err.Error())
	}
}
