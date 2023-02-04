package dto

import "github.com/mkeort/bath-hexagonal/domain/model"

type NewPoop struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" gorm:"default:null"`
	Latitude    string `json:"latitude" validate:"required"`
	Longitude   string `json:"Longitude" validate:"required"`
}

func (np *NewPoop) ToPoop(user model.User) *model.Poop {
	return &model.Poop{
		Title:       np.Title,
		Description: np.Description,
		Latitude:    np.Latitude,
		Longitude:   np.Longitude,
		User:        user,
	}
}
