package repository

import (
	"errors"

	"github.com/mkeort/bath-hexagonal/domain/model"
	"github.com/mkeort/bath-hexagonal/usecase/repository"
	"gorm.io/gorm"
)

type poopRepository struct {
	db *gorm.DB
}

func NewPoopRepository(db *gorm.DB) repository.PoopRepository {
	return &poopRepository{db}
}

func (pr *poopRepository) Create(p *model.Poop) (*model.Poop, error) {
	if err := pr.db.Create(p).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return p, nil
}

func (pr *poopRepository) GetAll() ([]model.Poop, error) {
	var poops []model.Poop
	if err := pr.db.Find(&poops).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return poops, nil
}

func (pr *poopRepository) GetMine(ui uint) ([]model.Poop, error) {
	var poops []model.Poop
	if err := pr.db.Where("user_id = ?", ui).Find(&poops).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return poops, nil
}
