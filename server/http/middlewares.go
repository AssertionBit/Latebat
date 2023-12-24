package http

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func DevCorsMiddleware(ctx *fiber.Ctx) error {
  ctx.Response().Header.Add("Access-Control-Allow-Origin", "http://localhost:5173")

  return ctx.Next()
}

func LoggingMiddleware(ctx *fiber.Ctx) error {
  logger.Info(
    "Retrieving new connection",
    zap.String("IP", ctx.IP()),
    zap.ByteString("X-Real-IP", ctx.Request().Header.Peek("X-Real-IP")),
    zap.ByteString("User-Agent", ctx.Request().Header.UserAgent()),
  )

  return ctx.Next()
}

