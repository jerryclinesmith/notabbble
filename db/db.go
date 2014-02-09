package db

import (
	"github.com/jerryclinesmith/notabbble/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

func Connect() gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_CONNECTION"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.Project{})
	// db.LogMode(true) // TODO: for dev only

	return db
}
