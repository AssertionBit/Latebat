package model

import (
	"os"

	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initSqliteDatabase() *gorm.DB {
  db, err := gorm.Open(sqlite.Open("./data.db"), nil)
  if err != nil {
    println(err)
    os.Exit(1)
  }

  db.AutoMigrate(
    &UserModel{}, 
    &DocumentModel{}, 
  )

  return db
}

func initPostgresDatabase() *gorm.DB {
  return nil 
}

func InitDatabase() *gorm.DB {
  return initSqliteDatabase()
}
