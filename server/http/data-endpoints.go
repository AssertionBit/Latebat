package http

import (
	"encoding/json"

	"fumine.ru/lebetat/model"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type authRequestT struct {
  Login    string `json:"login"`
  Password string `json:"password"`
}

type authResponseT struct {
  Name string `json:"name"`
  Type model.UserType `json:"type"`
}

func authPostEndpoint(ctx *fiber.Ctx) error {
  var authRequest authRequestT
  if err := json.Unmarshal(ctx.Body(), &authRequest); err != nil {
    logger.Warn(
      "Malformed data accepted",
      zap.String("IP", ctx.IP()),
    )

    return ctx.SendStatus(400)
  }

  // Check if user exists, throws error if false
  var possibleUser model.UserModel
  if err := db.Model(&model.UserModel{}).Where(&model.UserModel{ Login: authRequest.Login }).First(&possibleUser).Error; 
     err == nil {
     logger.Warn("User exists")
     return ctx.SendStatus(401)
  }

  possibleUser = model.UserModel{
    Name: authRequest.Login,
    Login: authRequest.Login,
    Password: authRequest.Password,
    Type: model.Default,
  }

  if err := db.Create(&possibleUser).Commit().Error; err != nil {
    logger.Error(
      "User creation failed at transaction step", 
      zap.String("error", err.Error()),
    )
  }
  logger.Info(
    "New user created", 
    zap.String("Login", possibleUser.Name),
    zap.Timep("When", &possibleUser.CreatedAt),
  )

  ctx.SendStatus(200)
  if res, err := json.Marshal(&authResponseT{
    Name: possibleUser.Name,
    Type: possibleUser.Type,
  }); err != nil {
    logger.Error(
      "Malformed body at creating response to end user",
      zap.String("Login", possibleUser.Name),
    )
    return ctx.SendStatus(500)

  } else {

    _ = ctx.SendStatus(200)
    return ctx.Send(res)
  }
}

func authGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func userGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func subjectGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func subjectPostEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func subjectsAllGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func subjectsDeleteEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func documentGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func documentGetAllEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func documentPostEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func documentUpdateEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func documentDeleteEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

