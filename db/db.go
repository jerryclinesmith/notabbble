package db

import (
	"github.com/codegangsta/martini"
	"github.com/jerryclinesmith/notabbble/models"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"log"
	"os"
)

func Connect() gorm.DB {
	url := os.Getenv("DATABASE_URL")
	werckerDbUrl := os.Getenv("WERCKER_POSTGRESQL_URL")
	if werckerDbUrl != "" {
		url = werckerDbUrl
	}

	connection, err := pq.ParseURL(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %s", err)
	}

	connection += " sslmode=disable" // TODO: Use require in production
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Fatalf("Failed to connecto to database: %s", err)
	}

	db.AutoMigrate(models.Project{})
	if martini.Env == martini.Dev {
		db.LogMode(true)
	}

	return db
}
