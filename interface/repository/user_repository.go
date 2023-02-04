package repository

import (
	"errors"

	"github.com/mkeort/bath-hexagonal/domain/model"
	"github.com/mkeort/bath-hexagonal/usecase/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Create(u *model.User) (*model.User, error) {
	if err := ur.db.Create(u).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) FindByEmail(e string) (*model.User, error) {
	var u model.User

	if err := ur.db.Where("email = ?", e).First(&u).Error; !errors.Is(err, nil) {
		return nil, errors.New("credentials don't match")
	}

	return &u, nil
}
