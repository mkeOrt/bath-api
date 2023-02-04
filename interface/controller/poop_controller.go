package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/usecase/interactor"
)

type poopController struct {
	poopInteractor interactor.PoopInteractor
}

type PoopController interface {
	Create(c *fiber.Ctx) error
}

func NewPoopController(us interactor.PoopInteractor) PoopController {
	return &poopController{us}
}

func (uc *poopController) Create(c *fiber.Ctx) error {
	return c.JSON(c.Locals("UserId"))
}
