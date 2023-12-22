package main

import (
	"fumine.ru/lebetat/http"
	"go.uber.org/zap"
)

func main() {
  logger, _ := zap.NewProduction()
  server := http.InitServer(logger)
  server.Listen(":8080")
}

