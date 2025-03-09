package database

import (
	"log"

	"streaming_demo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func InitDB() {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	err = db.AutoMigrate(&models.Record{})
	if err != nil {
		return
	}

	DB = Dbinstance{
		Db: db,
	}
}
