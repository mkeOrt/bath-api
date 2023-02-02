package interactor

import (
	"errors"

	"github.com/mkeort/bath-hexagonal/domain/dto"
	"github.com/mkeort/bath-hexagonal/domain/model"
	"github.com/mkeort/bath-hexagonal/usecase/presenter"
	"github.com/mkeort/bath-hexagonal/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
	DBRepository   repository.DBRepository
}

type UserInteractor interface {
	Create(u *model.User) (*dto.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter, d repository.DBRepository) UserInteractor {
	return &userInteractor{r, p, d}
}

func (us *userInteractor) Create(u *model.User) (*dto.User, error) {
	data, err := us.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		return us.UserRepository.Create(u)
	})

	user, ok := data.(*model.User)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(user), nil
}
