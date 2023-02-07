package interactor

import (
	"errors"
	"math"

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
	GetAll(pageSize, page int) (*dto.PaginatedPoopsWithUser, error)
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

func (pi *poopInteractor) GetAll(pageSize, page int) (*dto.PaginatedPoopsWithUser, error) {
	cPoops := make(chan *[]model.Poop)
	cCountPoops := make(chan *int64)
	go func() {
		poopsCount, err := pi.PoopRepository.GetAllCount()
		if err != nil {
			cPoops <- nil
		}
		cCountPoops <- poopsCount
	}()

	go func() {
		poops, err := pi.PoopRepository.GetAll(pageSize, page)
		if err != nil {
			cPoops <- nil
		}
		cPoops <- &poops
	}()

	poops := <-cPoops
	countPoops := <-cCountPoops

	if poops == nil || countPoops == nil {
		return nil, errors.New("error getting poops")
	}

	pp := dto.PaginatedPoopsWithUser{
		Count:       *countPoops,
		Page:        page,
		PageSize:    pageSize,
		PagesAmount: int(math.Ceil(float64(*countPoops) / float64(pageSize))),
		Poops:       pi.PoopPresenter.PoopsCreatedWithUser(*poops),
	}

	return &pp, nil
}

func (pi *poopInteractor) GetMine(ui uint) ([]dto.PoopCreated, error) {
	poops, err := pi.PoopRepository.GetMine(ui)

	if err != nil {
		return nil, errors.New("error getting poops")
	}

	return pi.PoopPresenter.PoopsCreated(poops), nil
}
