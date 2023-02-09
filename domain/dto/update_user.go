package dto

type UpdateUser struct {
	Name                 string `json:"name,omitempty" validate:"omitempty,min=1"`
	Lastname             string `json:"lastname,omitempty" validate:"omitempty,min=1"`
	Email                string `json:"email,omitempty" validate:"omitempty,email"`
	Password             string `json:"password,omitempty" validate:"omitempty,min=6"`
	PasswordConfirmation string `json:"password-confirmation" validate:"eqcsfield=Password"`
}
