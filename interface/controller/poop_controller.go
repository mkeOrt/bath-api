package controller

import (
	"errors"
	"strconv"

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
	GetMine(c *fiber.Ctx) error
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
	pageSize := c.Query("page_size", "10")
	intPageSize, err := strconv.Atoi(pageSize)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("page_size should be a valid number")
	}

	page := c.Query("page", "1")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("page should be a valid number")
	}

	poops, err := uc.poopInteractor.GetAll(intPageSize, intPage)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(poops)
}

func (uc *poopController) GetMine(c *fiber.Ctx) error {
	user := c.Locals("User").(model.User)
	poops, err := uc.poopInteractor.GetMine(user.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(poops)
}
