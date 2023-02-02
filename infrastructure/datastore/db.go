package datastore

import (
	"github.com/mkeort/bath-hexagonal/config"
	"github.com/mkeort/bath-hexagonal/domain/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.C.Database.Name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	return db
}
