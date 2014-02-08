package db

import (
	// "database/sql"
	"github.com/jerryclinesmith/notabbble/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Connect() gorm.DB {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := gorm.Open("postgres", "user=notabbble dbname=notabbble_dev sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.Project{})
	db.LogMode(true)

	return db
}
