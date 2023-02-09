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
	LogIn(l *dto.LogIn) (*dto.LoggedIn, error)
	GetMe(u *model.User) *dto.User
	Update(u *model.User, nu *dto.UpdateUser) (*dto.User, error)
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

func (us *userInteractor) LogIn(l *dto.LogIn) (*dto.LoggedIn, error) {
	user, err := us.UserRepository.FindByEmail(l.Email)
	if err != nil {
		return nil, err
	}
	isValidPassword := user.ValidatePassword(l.Password)
	if !isValidPassword {
		return nil, errors.New("credentials don't match")
	}

	return us.UserPresenter.ResponseLoggedIn(user)
}

func (ui *userInteractor) GetMe(u *model.User) *dto.User {
	return ui.UserPresenter.ResponseUser(u)
}

func (ui *userInteractor) Update(u *model.User, nu *dto.UpdateUser) (*dto.User, error) {
	newUser := model.User{
		Name:     nu.Name,
		Lastname: nu.Lastname,
		Email:    nu.Email,
	}
	if err := newUser.CryptPassword(nu.Password); err != nil {
		return nil, errors.New("error encrypting password")
	}

	data, err := ui.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		return ui.UserRepository.Update(u, &newUser)
	})

	if err != nil {
		return nil, err
	}

	user, ok := data.(*model.User)

	if !ok {
		return nil, errors.New("cast error")
	}

	return ui.UserPresenter.ResponseUser(user), nil
}
