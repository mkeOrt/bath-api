package interactor

import (
	"errors"

	"github.com/mkeort/bath-hexagonal/domain/dto"
	"github.com/mkeort/bath-hexagonal/domain/model"
	"github.com/mkeort/bath-hexagonal/usecase/presenter"
	"github.com/mkeort/bath-hexagonal/usecase/repository"
)

type poopInteractor struct {
	PoopRepository repository.PoopRepository
	PoopPresenter  presenter.PoopPresenter
	DBRepository   repository.DBRepository
}

type PoopInteractor interface {
	Create(p *model.Poop) (*dto.PoopCreated, error)
	GetAll() ([]dto.PoopCreatedWithUser, error)
	GetMine(ui uint) ([]dto.PoopCreated, error)
}

func NewPoopInteractor(r repository.PoopRepository, p presenter.PoopPresenter, d repository.DBRepository) PoopInteractor {
	return &poopInteractor{r, p, d}
}

func (pi *poopInteractor) Create(p *model.Poop) (*dto.PoopCreated, error) {
	data, err := pi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		return pi.PoopRepository.Create(p)
	})

	poop, ok := data.(*model.Poop)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}
	return pi.PoopPresenter.PoopCreated(poop), nil
}

func (pi *poopInteractor) GetAll() ([]dto.PoopCreatedWithUser, error) {
	poops, err := pi.PoopRepository.GetAll()

	if err != nil {
		return nil, errors.New("error getting poops")
	}

	return pi.PoopPresenter.PoopsCreatedWithUser(poops), nil
}

func (pi *poopInteractor) GetMine(ui uint) ([]dto.PoopCreated, error) {
	poops, err := pi.PoopRepository.GetMine(ui)

	if err != nil {
		return nil, errors.New("error getting poops")
	}

	return pi.PoopPresenter.PoopsCreated(poops), nil
}
