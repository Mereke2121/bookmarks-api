package main

import (
	"github.com/bookmarks-api/config"
	"github.com/bookmarks-api/pkg/handlers"
	"github.com/bookmarks-api/pkg/repository"
	"github.com/bookmarks-api/pkg/services"
	"github.com/bookmarks-api/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("load env variables; err: %s", err.Error())
	}

	conf, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("init config; err: %s", err.Error())
	}

	db, err := repository.Connect(conf)
	if err != nil {
		logrus.Fatalf("connect to postgres db; %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	srv := new(server.Server)
	if err = srv.Run(conf.Port, handler.InitRoutes()); err != nil {
		logrus.Fatalf("try to run server; err: %s", err.Error())
	}
}
