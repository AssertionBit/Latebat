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
  var possibleSubject model.SubjectModel

  id, err := ctx.ParamsInt("id", -1)
  if err != nil {
    logger.Warn(
      "Malformed path param (possibly missing)",
      zap.String("Request-type", "retrieve single subject"),
    )
  }

  // Seaching for single subject, error means not exists
  if err := db.Model(&model.SubjectModel{}).Where("id = ?", id).First(&possibleSubject).Error; err != nil {
    logger.Info(
      "User not found with such id", 
      zap.Int("Subject-Id", id),
      zap.String("Request-type", "retrieve single subject"),
    )

    return ctx.SendStatus(404)
  }

  // Processing for response
  data := subjectResponseFullT{
    Name: "",
    Documents: []documentResponseT{},
  }

  for _, doc := range possibleSubject.DocumentModels {
    data.Documents = append(data.Documents, documentResponseT{
      Id: doc.ID,
      BacktrackID: doc.BacktrackId,
      Name: doc.Name,
      Status: doc.Status,
    })
  }

  if res, err := json.Marshal(&data); err != nil {
    return ctx.SendStatus(500)
  } else {
    _ = ctx.SendStatus(200)
    return ctx.Send(res)
  }
}

/// subjectsAllGetEndpoint function retrieves all subjects
/// and returns array of their files, and them selfs. Function serves
/// as just mark that document will be processed. For general data upload,
/// use `/api/v1/upload` endpoint or uploadDocumentEndpoint function
///
/// Params:
///   ctx - Context for request/response from fiber.
func subjectsAllGetEndpoint(ctx *fiber.Ctx) error {
  var subjects []model.SubjectModel = make([]model.SubjectModel, 0)

  if err := db.Model(&model.SubjectModel{}).Find(&subjects).Error; err != nil {
    logger.Info("No users found, returning empty array")
    if res, err := json.Marshal(&subjects); err != nil {
      return ctx.SendStatus(500)
    } else {
      return ctx.Send(res)
    }
  }

  var subjectRespo []subjectResponseT = make([]subjectResponseT, 0)

  logger.Info("Processing all subjects", zap.Int("Length", len(subjects)))
  for _, subj := range subjects {
    docId := make([]uint, 0)
    for _, doc := range subj.DocumentModels {
      docId = append(docId, doc.ID)
    }

    subjectRespo = append(subjectRespo, subjectResponseT{
      Name: "",
      DocumentID: docId,
    })
  }

  logger.Info(
    "Processed data succesfully", 
    zap.String("Request-type", "Retreiving all subjects"),
  )
  if res, err := json.Marshal(subjectRespo); err != nil {
    return ctx.SendStatus(500)
  } else {
    _ = ctx.SendStatus(200)
    return ctx.Send(res)
  }
}

func subjectPostEndpoint(ctx *fiber.Ctx) error {
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

/// documentGeneralPostEndpoint function which represents general data
/// reciving point for this service. Goal is to save and handle the data from
/// end user.
func documentGeneralPostEndpoint(ctx *fiber.Ctx) error {
  return ctx.SendStatus(200)
}

