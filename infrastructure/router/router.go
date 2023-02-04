package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/interface/controller"
)

func NewRouter(r fiber.Router, c controller.AppController) {
	middlewares := c.Middleware.GetAuthMiddlewares()

	NewUserRouter(r.Group("/users"), c)
	NewPoopRouter(r.Group("/poops", middlewares.RequiredAuth), c)
}
