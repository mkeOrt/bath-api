package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/interface/controller"
)

func NewRouter(a fiber.Router, c controller.AppController) {
	users := a.Group("/users")

	users.Post("sign-up", c.User.SignUp)
}
