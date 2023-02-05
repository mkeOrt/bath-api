package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/domain/dto"
	"github.com/mkeort/bath-hexagonal/domain/model"
	"github.com/mkeort/bath-hexagonal/usecase/interactor"
)

type poopController struct {
	poopInteractor interactor.PoopInteractor
}

type PoopController interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

func NewPoopController(us interactor.PoopInteractor) PoopController {
	return &poopController{us}
}

func (uc *poopController) Create(c *fiber.Ctx) error {
	var np dto.NewPoop
	if err := c.BodyParser(&np); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errs := dto.ValidateStruct(np); errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errs)

	}

	user := c.Locals("User").(model.User)

	cp, err := uc.poopInteractor.Create(np.ToPoop(model.User(user)))
	if !errors.Is(err, nil) {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.JSON(cp)
}

func (uc *poopController) GetAll(c *fiber.Ctx) error {
	poops, err := uc.poopInteractor.GetAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(poops)
}
