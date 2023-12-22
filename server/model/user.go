package model

import "gorm.io/gorm"

type UserType string

const (
  Default       UserType = "default"
  Administrator UserType = "admin"
)

type UserModel struct {
  gorm.Model

  Name     string   `json:"name"`
  Login    string
  Password string
  Type     UserType `json:"type"`
}

