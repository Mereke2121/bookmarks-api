package main

import (
	"fmt"
	"github.com/bookmarks-api/pkg/repository"
	"github.com/bookmarks-api/pkg/services"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"

	"github.com/bookmarks-api/pkg/handlers"
	"github.com/bookmarks-api/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("load env variables", err)
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=postgres dbname=bookmarks sslmode=disable password=%s", os.Getenv("DB_PASSWORD")))
	if err != nil {
		log.Fatal("open postgres database", err)
	}

	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	srv := new(server.Server)
	if err = srv.Run(":8000", handler.InitRoutes()); err != nil {
		log.Fatal("try to run server", err)
	}
}
