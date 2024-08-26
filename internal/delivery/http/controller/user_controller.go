package controller

import (
	"edot-monorepo/services/user-service/internal/model"
	"edot-monorepo/services/user-service/internal/usecase"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	UserRegisterUseCase *usecase.UserRegisterUseCase
	UserLoginUseCase    *usecase.UserLoginUseCase
	Log                 *logrus.Logger
	Validate            *validator.Validate
}

func NewUserController(userRegisterUseCase *usecase.UserRegisterUseCase, userLoginUseCase *usecase.UserLoginUseCase, log *logrus.Logger, validate *validator.Validate) *UserController {
	return &UserController{
		UserRegisterUseCase: userRegisterUseCase,
		UserLoginUseCase:    userLoginUseCase,
		Log:                 log,
		Validate:            validate,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.UserRegisterRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UserRegisterUseCase.Exec(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.UserLoginRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UserLoginUseCase.Exec(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AuthResponse]{
		Data: response,
	})

}
