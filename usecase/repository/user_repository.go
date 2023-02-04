package repository

import (
	"github.com/mkeort/bath-hexagonal/domain/model"
)

type UserRepository interface {
	GetById(uid int) (*model.User, error)
	Create(u *model.User) (*model.User, error)
	FindByEmail(e string) (*model.User, error)
}
