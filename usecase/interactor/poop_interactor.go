package interactor

import (
	"github.com/mkeort/bath-hexagonal/usecase/presenter"
	"github.com/mkeort/bath-hexagonal/usecase/repository"
)

type poopInteractor struct {
	PoopRepository repository.PoopRepository
	PoopPresenter  presenter.PoopPresenter
	DBRepository   repository.DBRepository
}

type PoopInteractor interface {
}

func NewPoopInteractor(r repository.PoopRepository, p presenter.PoopPresenter, d repository.DBRepository) PoopInteractor {
	return &poopInteractor{r, p, d}
}
