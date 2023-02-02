package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/domain/dto"
	"github.com/mkeort/bath-hexagonal/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	SignUp(c *fiber.Ctx) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) SignUp(c *fiber.Ctx) error {
	var signUp dto.SignUp
	if err := c.BodyParser(&signUp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errs := dto.ValidateStruct(signUp)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errs)

	}

	u, err := uc.userInteractor.Create(signUp.ToUser())
	if !errors.Is(err, nil) {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}
