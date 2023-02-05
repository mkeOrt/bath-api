package dto

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mkeort/bath-hexagonal/config"
	"github.com/mkeort/bath-hexagonal/domain/model"
)

type LoggedIn struct {
	Token string `json:"token"`
}

func NewLoggedIn(u *model.User) (*LoggedIn, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"nbf":     jwt.NewNumericDate(time.Now()),
		"iat":     jwt.NewNumericDate(time.Now()),
		"user_id": u.ID,
	})

	ss, err := token.SignedString([]byte(config.C.JWT.SecretKey))
	if err != nil {
		return nil, err
	}

	return &LoggedIn{
		Token: fmt.Sprintf("%v", ss),
	}, nil
}
