package config

import (
	"edot-monorepo/services/user-service/internal/delivery/http/controller"
	"edot-monorepo/services/user-service/internal/delivery/http/route"
	"edot-monorepo/services/user-service/internal/gateway/http"
	repository "edot-monorepo/services/user-service/internal/repository/gorm"
	"edot-monorepo/services/user-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB         *gorm.DB
	App        *fiber.App
	Log        *logrus.Logger
	Validate   *validator.Validate
	Config     *viper.Viper
	HttpClient *resty.Client
}

func Bootstrap(config *BootstrapConfig) {

	kongClient := http.NewKongClient(config.HttpClient, config.Config.GetString("kong.jwt-consumer"))

	userRepository := repository.NewUserRepository(config.Log)
	userBaseUseCase := usecase.NewUserUseCase(config.DB, config.Log, userRepository, config.Validate, kongClient)
	userLoginUseCase := usecase.NewUserLoginUseCase(userBaseUseCase)
	userRegisterUseCase := usecase.NewUserRegisterUseCase(userBaseUseCase)
	userController := controller.NewUserController(userRegisterUseCase, userLoginUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}

	routeConfig.Setup()
}
