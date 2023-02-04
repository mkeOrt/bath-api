package presenter

import (
	"github.com/mkeort/bath-hexagonal/domain/dto"
	"github.com/mkeort/bath-hexagonal/domain/model"
)

type UserPresenter interface {
	ResponseUser(u *model.User) *dto.User
	ResponseLoggedIn(u *model.User) (*dto.LoggedIn, error)
}
