package usecase

import (
	repository "edot-monorepo/services/user-service/internal/repository/gorm"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserBaseUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	UserRepository *repository.UserRepository
	Validate       *validator.Validate
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, userRepo *repository.UserRepository, validate *validator.Validate) *UserBaseUseCase {
	return &UserBaseUseCase{
		DB:             db,
		Log:            log,
		UserRepository: userRepo,
		Validate:       validate,
	}
}
