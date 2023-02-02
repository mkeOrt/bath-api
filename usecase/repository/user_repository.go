package repository

import (
	"github.com/mkeort/bath-hexagonal/domain/model"
)

type UserRepository interface {
	Create(u *model.User) (*model.User, error)
}
