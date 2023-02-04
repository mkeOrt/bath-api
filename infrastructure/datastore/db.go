package datastore

import (
	"fmt"
	"os"

	"github.com/mkeort/bath-hexagonal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=bath-database user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Mexico_City",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Poop{})

	return db
}
