package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/interface/controller"
)

func NewUserRouter(r fiber.Router, c controller.AppController) {
	middlewares := c.Middleware.GetAuthMiddlewares()

	r.Get("me", middlewares.RequiredAuth, c.User.GetUser)
	r.Post("sign-up", c.User.SignUp)
	r.Post("log-in", c.User.LogIn)
}
