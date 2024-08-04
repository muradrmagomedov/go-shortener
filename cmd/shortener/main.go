package main

import (
	"github.com/mattn/go-colorable"
	"github.com/muradrmagomedov/go-shortener/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	initLogger()
	server := handler.NewServer()
	server.Run(":8080")
}

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
}
