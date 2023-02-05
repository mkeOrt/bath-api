package presenter

import (
	"github.com/mkeort/bath-hexagonal/domain/dto"
	"github.com/mkeort/bath-hexagonal/domain/model"
)

type PoopPresenter interface {
	PoopCreated(p *model.Poop) *dto.PoopCreated
	PoopsCreated(poops []model.Poop) []dto.PoopCreated
}
