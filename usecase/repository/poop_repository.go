package repository

import (
	"github.com/mkeort/bath-hexagonal/domain/model"
)

type PoopRepository interface {
	Create(u *model.Poop) (*model.Poop, error)
	GetAll(pageSize, page int) ([]model.Poop, error)
	GetMine(ui uint) ([]model.Poop, error)
	GetAllCount() (*int64, error)
}
