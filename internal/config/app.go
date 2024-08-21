package config

import (
	"edot-monorepo/services/user-service/internal/delivery/http/controller"
	"edot-monorepo/services/user-service/internal/delivery/http/route"
	repository "edot-monorepo/services/user-service/internal/repository/gorm"
	"edot-monorepo/services/user-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {

	userRepository := repository.NewUserRepository(config.Log)
	userBaseUseCase := usecase.NewUserUseCase(config.DB, config.Log, userRepository, config.Validate)
	userLoginUseCase := usecase.NewUserLoginUseCase(userBaseUseCase)
	userRegisterUseCase := usecase.NewUserRegisterUseCase(userBaseUseCase)
	userController := controller.NewUserController(userRegisterUseCase, userLoginUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}

	routeConfig.Setup()
}
