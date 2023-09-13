package main

import (
	"Forum"
	"Forum/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	server := new(Forum.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while starting server")
	}
}
