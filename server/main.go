package main

import (
	"fumine.ru/lebetat/http"
	"fumine.ru/lebetat/model"
	"go.uber.org/zap"
)

func main() {
  logger, _ := zap.NewProduction()
  db := model.InitDatabase()

  server := http.InitServer(logger, db)
  server.Listen(":8080")
}

