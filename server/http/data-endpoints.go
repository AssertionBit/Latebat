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

type documentResponseT struct {
  Id uint `json:"id"`
  Name string `json:"name"`
  Status model.DocumentT `json:"status"`
  BacktrackID string `json:"backtrack-id"`
}

type subjectResponseT struct {
  Name string `json:"name"`
  DocumentID []uint `json:"documents-id"`
}

type subjectResponseFullT struct {
  Name string `json:"name"`
  Documents []documentResponseT `json:"documents"`
}

/// authPostEndpoint takes raw JSON type and prcesses it 
/// for creation of user. If user with such credentials
/// exists, then user will not be created.
///
/// Params:
///    ctx - Context of request/response from fiber
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

  // All is going okay
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

/// subjectGetEndpoint function retrieves information single
/// subject and returns all information about his documents.
/// 
/// Param:
///   id  - Id of subject
///   ctx - request/response context from fiber
///
/// Return:
///   JSON object when possible 
func subjectGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(404)
}

/// subjectsAllGetEndpoint function retrieves all subjects
/// and returns array of their files, and them selfs. Function serves
/// as just mark that document will be processed. For general data upload,
/// use `/api/v1/upload` endpoint or uploadDocumentEndpoint function
///
/// Params:
///   ctx - Context for request/response from fiber.
func subjectsAllGetEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(404)
}

func subjectPostEndpoint(ctx *fiber.Ctx) error {
  // ctx.Accepts("image/png")
  // ctx.Accepts("image/jpeg")
  // ctx.Accepts("application/pdf")
  ctx.AcceptsEncodings("utf-8")

  logger.Info(
    "Accepting new multipart request",
  )
  form, err := ctx.MultipartForm()

  if err != nil {
    logger.Error(
      "Request failed to be processed",
      zap.Error(err),
    )
    return ctx.SendStatus(500)
  }

  for _, formHeader := range form.File {
  
    for _, fileHeader := range formHeader {
      fileRaw, err := fileHeader.Open()
      if err != nil {
        logger.Warn(
          "Failed to open file",
        zap.String("File-Type", fileHeader.Header.Get("Content-Type")),
          zap.String("File-Name", fileHeader.Filename),
        )
        continue
      }
      var fileCont []byte
      if _, err := fileRaw.Read(fileCont); err != nil {
        logger.Warn(
          "Failed to open file",
          zap.String("File-Type", fileHeader.Header.Get("Content-Type")),
          zap.String("File-Name", fileHeader.Filename),
        )
        continue
      }

      file := model.DocumentModel{
    Status: model.Accepted,
    Format: fileHeader.Header.Get("Content-Type"),
    Content: fileCont,
  }

  // IF error, data will not be saved
  if err := db.Create(&file).Commit().Error; err != nil {
    logger.Warn(
      "File failed to be saved",
      zap.ByteString("Content-Type", ctx.Request().Header.ContentType()),
      zap.Error(err),
    )

    return ctx.SendStatus(500)
  }

  logger.Info(
    "File accepted and will be processed",
    zap.String("Content-Type", file.Format),
    zap.Int("Size", int(len(file.Content))),
  )
}
  }

  return ctx.SendStatus(200)
}

func subjectsDeleteEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func documentGetEndpoint(ctx *fiber.Ctx) error {
  id, err := ctx.ParamsInt("docId", -1)
  logger.Info(
    "Sending image",
    zap.Int("ID", id),
  )
  if err != nil {
    logger.Warn("Error")
    return ctx.SendStatus(400)
  }

  var doc model.DocumentModel

  if err := db.Model(&model.DocumentModel{}).Where("id = ?", id).First(&doc).Error; err != nil {
    return ctx.SendStatus(404)
  }
 
  logger.Info("Sending file", zap.Int("ID", id))
  ctx.Response().Header.Set("Content-Type", doc.Format)
  return ctx.Send(doc.Content)
}

func documentGetAllEndpoint(ctx *fiber.Ctx) error {
  data := make([]fiber.Map, 0)

  var dbData []model.DocumentModel
  if err := db.Model(&model.DocumentModel{}).Find(&dbData).Error; err != nil {
    logger.Warn(
      "No data found in database, nothing will be returnder",
    )
  } else {
    for _, mod := range dbData {
      data = append(data, fiber.Map{
        "id": mod.ID,
        "name": mod.Name,
        "status": mod.Status,
        "format": mod.Format,
      })
    }
  }

  if data, err := json.Marshal(data); err != nil {
    return ctx.SendStatus(500)
  } else {
    return ctx.Send(data)
  }
}

func documentPostEndpoint(ctx *fiber.Ctx) error {
  ctx.Accepts("multipart/form-data")

  logger.Info(
    "Accepting new connection",
  )

  file, err := ctx.FormFile("file")

  if err != nil {
    logger.Error(
      "Error handling file from request",
      zap.Error(err),
    )
    return ctx.SendStatus(500)
  }

  rawFile, err := file.Open()
  if err != nil {
    logger.Error(
      "Error handling file from request",
      zap.Error(err),
    )
    return ctx.SendStatus(500)
  }

  body := make([]byte, 150000)
  if n, err := rawFile.Read(body); err != nil {
      logger.Error(
        "Error handling file from request",
        zap.Error(err),
      )
      return ctx.SendStatus(500)
  } else {
    logger.Info(
      "Info attempting saving file to database",
      zap.Int("Size", n),
      zap.String("Content-Type", file.Header.Get("Content-Type")),
      zap.String("File-Name", file.Filename),
    )
  }

  fileDto := model.DocumentModel{
    Name: file.Filename,
    Status: model.Accepted,
    Format: file.Header.Get("Content-Type"),
    Content: body,
  }

  if err := db.Create(&fileDto).Error; err != nil {
    logger.Warn(
      "Error saving file in database",
      zap.String("Content-Type", fileDto.Format),
      zap.Error(err),
    )

    return ctx.SendStatus(500)
  }

  return ctx.RedirectBack("/")
}

func documentUpdateEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

func documentDeleteEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

/// documentGeneralPostEndpoint function which represents general data
/// reciving point for this service. Goal is to save and handle the data from
/// end user.
func documentGeneralPostEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

