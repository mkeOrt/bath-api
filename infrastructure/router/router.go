package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/interface/controller"
	"github.com/mkeort/bath-hexagonal/interface/middleware"
)

func NewRouter(r fiber.Router, c controller.AppController) {
	middlewares := middleware.NewAppMiddleware()
	NewUserRouter(r.Group("/users"), c)
	NewPoopRouter(r.Group("/poops", middlewares.Auth.RequiredAuth), c)
}
