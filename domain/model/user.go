package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Poops    []Poop
}

func NewUser(n, l, e, p string) (*User, error) {
	hash, err := GenerateFromPassword(p)
	if err != nil {
		return nil, fmt.Errorf("error encrypting password")
	}
	return &User{
		Name:     n,
		Lastname: l,
		Email:    e,
		Password: hash,
	}, nil
}

func (u *User) ValidatePassword(p string) bool {
	return CompareHashAndPassword([]byte(u.Password), []byte(p))
}
