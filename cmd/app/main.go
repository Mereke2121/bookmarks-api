package main

import (
	"log"

	"github.com/bookmarks-api/pkg/handler"
	"github.com/bookmarks-api/server"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run(":8000", handlers.InitRoutes()); err != nil {
		log.Fatal("try to run server", err)
	}
}
