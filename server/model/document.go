package model

import "gorm.io/gorm"

type DocumentT string

const (
  Accepted   DocumentT = "accepted"
  Processing DocumentT = "processing"
  Processed  DocumentT = "processed"
)

/// Document - structure which represents document from one person
type DocumentModel struct {
  gorm.Model

  Name        string    `json:"name"`
  Status      DocumentT `json:"status"`
  Format      string    `json:"format"`
  Content     []byte
  BacktrackId string    `json:"backtrack-id"`
  SubjectModelID   uint
}

