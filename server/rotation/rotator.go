package rotation

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"fumine.ru/lebetat/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Rotator struct {
  handling []model.DocumentModel
  l *zap.Logger
  db *gorm.DB
}

func InitRotator(logger *zap.Logger, database *gorm.DB) (*Rotator, error) {
  r := &Rotator{
    handling: make([]model.DocumentModel, 0),
    l: logger,
    db: database,
  }

  var handleStart []model.DocumentModel
  if err := r.db.Model(&model.DocumentModel{}).Where(&model.DocumentModel{Status: model.Accepted}).Limit(10).Find(&handleStart).Error; err != nil {
    return nil, errors.New("Failed to retrieve data from database, possibly missing")
  }

  r.handling = handleStart

  return r, nil
}

func (rotator *Rotator) Exec() {
  go func () {
    for {
      rotator.l.Info("Executing new cycle", zap.Int("Items in queue", len(rotator.handling)))
      if len(rotator.handling) == 0 {
        rotator.l.Info("No files in queue. Waiting for new files")
        time.Sleep(time.Second * 10)
      }
      rotator.execCycle()

      var rotatedModels []model.DocumentModel
      if err := rotator.db.Where(&model.DocumentModel{Status: model.Accepted}).Limit(10).Find(&rotatedModels).Error; err != nil {
        rotator.l.Warn(
          "Possibly zero models left",
          zap.Error(err),
        )
        continue
      }

      rotator.handling = rotatedModels
    }
  }()
}

/// TODO: Extract processing to another method
func (rotator *Rotator) execCycle() {
  for _, f := range rotator.handling {
    rotator.l.Info("Processing file for anonimizing", zap.String("Name", f.Name))
    f.Status = model.Processing
    
    if err := rotator.db.Save(&f).Error; err != nil {
      rotator.l.Error(
        "Database doesn't saved model. Is database working normally?",
        zap.Uint("File-ID", f.ID),
        zap.Error(err),
      )
    }

    fileFormat := strings.Split(f.Format, "/")[1]
    os.WriteFile(fmt.Sprintf("/tmp/image.%s", fileFormat), f.Content, 0644)
    cmd := exec.Command("python3", "-m", "letebat", "process", fmt.Sprintf("/tmp/image.%s", fileFormat))
    if err := cmd.Run(); err != nil {
      rotator.l.Error(
        "For some reason failed to process file, file will excluded from cycle",
        zap.Uint("File-ID", f.ID),
        zap.Error(err),
      )

      f.Status = model.Processed

      if err := rotator.db.Save(&f).Error; err != nil {
        rotator.l.Error(
          "Database doesn't saved model. Is database working normally?",
          zap.Uint("File-ID", f.ID),
          zap.Error(err),
        )
      }

    } else {
      if data, err := os.ReadFile(fmt.Sprintf("/tmp/image.%s")); err != nil {
        rotator.l.Warn(
          "Failed to read file, other processes could not read it, so problem in FS",
          zap.Uint("File-ID", f.ID),
          zap.Error(err),
        )
      } else {
        f.Status = model.Processed
        f.Content = data
      
        if err := rotator.db.Model(&model.DocumentModel{}).Save(&f).Error; err != nil {
          rotator.l.Error(
            "Database doesn't saved model. Is database working normally?",
            zap.Uint("File-ID", f.ID),
            zap.Error(err),
          )
        }
      }
    }
    os.Remove(fmt.Sprintf("/tmp/image.%s", fileFormat))
  }
}

