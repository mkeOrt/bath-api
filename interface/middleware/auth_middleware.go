package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type authMiddleware struct {
}

type AuthMiddleware interface {
	RequiredAuth(c *fiber.Ctx) error
}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

func (uc *authMiddleware) RequiredAuth(c *fiber.Ctx) error {
	tokenString := c.Get("authorization")
	tss := strings.Split(tokenString, " ")
	if tss[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid token")
	}

	token, err := jwt.Parse(tss[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("UserId", fmt.Sprintf("%v", claims["user_id"]))
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON("Error reading token")
	}

	return c.Next()
}
