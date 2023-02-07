package presenter

import (
	"github.com/mkeort/bath-hexagonal/domain/dto"
	"github.com/mkeort/bath-hexagonal/domain/model"
	"github.com/mkeort/bath-hexagonal/usecase/presenter"
)

type poopPresenter struct{}

func NewPoopPresenter() presenter.PoopPresenter {
	return &poopPresenter{}
}

func (*poopPresenter) PoopCreated(p *model.Poop) *dto.PoopCreated {
	return &dto.PoopCreated{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Latitude:    p.Latitude,
		Longitude:   p.Longitude,
	}
}

func (*poopPresenter) PoopsCreated(poops []model.Poop) []dto.PoopCreated {
	var pp []dto.PoopCreated

	for _, p := range poops {
		pp = append(pp, dto.PoopCreated{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			Latitude:    p.Latitude,
			Longitude:   p.Longitude,
			LastUpdate:  p.UpdatedAt,
		})
	}

	return pp
}

func (*poopPresenter) PoopsCreatedWithUser(poops []model.Poop) []dto.PoopCreatedWithUser {
	var pp []dto.PoopCreatedWithUser

	for _, p := range poops {
		pp = append(pp, dto.PoopCreatedWithUser{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			Latitude:    p.Latitude,
			Longitude:   p.Longitude,
			LastUpdate:  p.UpdatedAt,
			User: dto.User{
				ID:       p.User.ID,
				Name:     p.User.Name,
				Lastname: p.User.Lastname,
				Email:    p.User.Email,
			},
		})
	}

	return pp
}
