package main

import (
	"github.com/mattn/go-colorable"
	"github.com/muradrmagomedov/go-shortener/internal/handler"
	"github.com/muradrmagomedov/go-shortener/internal/repository"
	"github.com/muradrmagomedov/go-shortener/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	initLogger()
	//TODO Вынести path в конфиг
	db, err := repository.NewSqliteConnection("./db/shortener.db")
	if err != nil {
		logrus.Fatal(err)
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	server := handler.NewHandler(service)
	//TODO Вынести port в конфиг
	server.Run(":8080")
}

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
}
