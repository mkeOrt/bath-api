package dto

import "github.com/mkeort/bath-hexagonal/domain/model"

type SignUp struct {
	Name                 string `json:"name" validate:"required"`
	Lastname             string `json:"lastname" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password-confirmation" validate:"required,eqcsfield=Password"`
}

func (s *SignUp) ToUser() *model.User {
	return &model.User{
		Name:     s.Name,
		Lastname: s.Lastname,
		Email:    s.Email,
		Password: s.Password,
	}
}
