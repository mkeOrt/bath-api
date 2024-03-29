package dto

import (
	"github.com/mkeort/bath-hexagonal/domain/model"
)

type SignUp struct {
	Name                 string `json:"name" validate:"required,min=1"`
	Lastname             string `json:"lastname" validate:"required,min=1"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password-confirmation" validate:"required,eqcsfield=Password"`
}

func (s *SignUp) ToUser() (*model.User, error) {
	return model.NewUser(s.Name, s.Lastname, s.Email, s.Password)
}
