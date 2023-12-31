package http

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var db *gorm.DB
var logger *zap.Logger


func InitServer(l *zap.Logger, database *gorm.DB) (*fiber.App) {
  logger = l
  db = database
  app := fiber.New()
  logger.Info("Initiating server")

  // 0. Using middlewares for logging
  app.Use(LoggingMiddleware, DevCorsMiddleware)

  // 1. Registering open endpoints
  app.Get("/api/v1/status", statusGetEndpoint)
  app.Get("/api/v1/auth", authGetEndpoint)
  app.Post("/api/v1/auth", authPostEndpoint)

  // 2.2 Subject documents
  app.Get("/api/v1/docs/", documentGetAllEndpoint)
  app.Get("/api/v1/docs/:docId", documentGetEndpoint)
  app.Post("/api/v1/docs/", documentPostEndpoint)
  app.Delete("/api/v1/docs/", documentDeleteEndpoint)
  app.Post("/api/v1//docs/:docId", documentUpdateEndpoint)

  // 2.3 Consumer endpoints
  app.Get("/api/v1/consumer", consumerAllGetEndpoint)
  app.Get("/api/v1/consumer/:subId", consumerGetEndpoint)

  // 2.4 General endpoints
  app.Post("/api/v1/upload", documentGeneralPostEndpoint)

  return app
}

