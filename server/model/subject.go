package model

import "gorm.io/gorm"


/// Subject - person, who uses service
type SubjectModel struct {
  gorm.Model

  DocumentModels []DocumentModel `json:"documents" gorm:"column:documents"`
}

