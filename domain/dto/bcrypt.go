package dto

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) (string, error) {
	cost, err := strconv.Atoi(os.Getenv("B_CRYPT_COST"))
	if err != nil {
		return "", fmt.Errorf("error encrypting data")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hash), err
}
