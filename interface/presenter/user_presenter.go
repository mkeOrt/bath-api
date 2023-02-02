package presenter

import (
	"github.com/mkeort/bath-hexagonal/domain/dto"
	"github.com/mkeort/bath-hexagonal/domain/model"
	"github.com/mkeort/bath-hexagonal/usecase/presenter"
)

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (*userPresenter) ResponseUser(us *model.User) *dto.User {
	return &dto.User{
		Name:     us.Name,
		Lastname: us.Lastname,
		Email:    us.Email,
	}
}
