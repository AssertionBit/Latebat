package http

import "github.com/gofiber/fiber/v2"

func consumerAllGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func consumerGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

