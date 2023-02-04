package repository

import (
	"github.com/mkeort/bath-hexagonal/usecase/repository"
	"gorm.io/gorm"
)

type poopRepository struct {
	db *gorm.DB
}

func NewPoopRepository(db *gorm.DB) repository.PoopRepository {
	return &poopRepository{db}
}
