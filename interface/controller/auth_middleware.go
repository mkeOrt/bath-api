package controller

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mkeort/bath-hexagonal/domain/model"
	"gorm.io/gorm"
)

type authMiddleware struct {
	DB *gorm.DB
}

type AuthMiddleware interface {
	RequiredAuth(c *fiber.Ctx) error
}

func NewAuthMiddleware(db *gorm.DB) AuthMiddleware {
	return &authMiddleware{
		DB: db,
	}
}

func (uc *authMiddleware) RequiredAuth(c *fiber.Ctx) error {
	tokenString := c.Get("authorization")
	tss := strings.Split(tokenString, " ")
	if tss[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid token")
	}

	token, err := jwt.Parse(tss[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user model.User
		if err := uc.DB.First(&user, fmt.Sprintf("%v", claims["user_id"])).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON("Error reading user")
		}
		fmt.Println(user)
		c.Locals("User", user)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON("Error reading token")
	}

	return c.Next()
}
