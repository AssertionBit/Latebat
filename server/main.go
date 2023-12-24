package main

import (
	"fumine.ru/lebetat/http"
	"fumine.ru/lebetat/model"
	"fumine.ru/lebetat/rotation"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	db := model.InitDatabase()
  rotator, err := rotation.InitRotator(logger, db)

  if err != nil {
    logger.Fatal(
      "Failed to start rotation process",
      zap.Error(err),
    )
  }

	server := http.InitServer(logger, db)
	rotator.Exec()
  server.Listen(":8080")
}
