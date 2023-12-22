package http

import "github.com/gofiber/fiber/v2"

type statusResponse struct {
  Status string `json:"status"`
}

func statusGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.JSON(statusResponse{Status: "normal"})
}

