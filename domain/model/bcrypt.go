package model

import (
	"github.com/mkeort/bath-hexagonal/config"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), config.C.BCrypt.Cost)
	return string(hash), err
}

func CompareHashAndPassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
