package main

import (
	"fumine.ru/lebetat/http"
	"fumine.ru/lebetat/model"
	"go.uber.org/zap"
)

func main() {
  logger, _ := zap.NewProduction()
  _ = model.InitDatabase()

  server := http.InitServer(logger)
  server.Listen(":8080")
}

