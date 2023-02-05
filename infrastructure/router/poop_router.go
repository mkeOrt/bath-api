package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/interface/controller"
)

func NewPoopRouter(r fiber.Router, c controller.AppController) {
	r.Post("create", c.Poop.Create)
	r.Get("get-all", c.Poop.GetAll)
}
