package repository

import (
	"errors"

	"github.com/mkeort/bath-hexagonal/domain/model"
	"github.com/mkeort/bath-hexagonal/infrastructure/datastore"
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

func (pr *poopRepository) GetAll(pageSize, page int) ([]model.Poop, error) {
	var poops []model.Poop

	if err := pr.db.Scopes(datastore.Paginate(pageSize, page)).Joins("User").Find(&poops).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return poops, nil
}

func (pr *poopRepository) GetAllCount() (*int64, error) {
	var count int64
	if err := pr.db.Model(&model.Poop{}).Count(&count).Error; err != nil {
		return nil, errors.New("error getting poops amount")
	}
	return &count, nil
}

func (pr *poopRepository) GetMine(ui uint) ([]model.Poop, error) {
	var poops []model.Poop
	if err := pr.db.Where("user_id = ?", ui).Find(&poops).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return poops, nil
}
