package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/domain/dto"
	"github.com/mkeort/bath-hexagonal/domain/model"
	"github.com/mkeort/bath-hexagonal/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	SignUp(c *fiber.Ctx) error
	LogIn(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
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

	user, err := signUp.ToUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	u, err := uc.userInteractor.Create(user)
	if !errors.Is(err, nil) {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}

func (uc *userController) LogIn(c *fiber.Ctx) error {
	var logIn dto.LogIn
	if err := c.BodyParser(&logIn); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errs := dto.ValidateStruct(logIn)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errs)

	}

	user, err := uc.userInteractor.LogIn(&logIn)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.JSON(user)
}

func (uc *userController) GetUser(c *fiber.Ctx) error {
	user := c.Locals("User").(model.User)
	return c.JSON(uc.userInteractor.GetMe(&user))
}
