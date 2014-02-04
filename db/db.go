package db

import (
	// "database/sql"
	"github.com/jinzhu/gorm"
	"github.com/jerryclinesmith/notabbble/models"
	_ "github.com/lib/pq"
)

func Connect() (gorm.DB) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := gorm.Open("postgres", "user=notabbble dbname=notabbble_dev sslmode=disable")
	if err != nil {
		panic(err)
	}
	
	db.AutoMigrate(models.Project{})
	
	return db
}
