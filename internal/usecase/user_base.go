package usecase

import (
	"edot-monorepo/services/user-service/internal/gateway/http"
	"edot-monorepo/services/user-service/internal/model"
	repository "edot-monorepo/services/user-service/internal/repository/gorm"

	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserBaseUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	UserRepository *repository.UserRepository
	Validate       *validator.Validate
	KongClient     *http.KongClient
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, userRepo *repository.UserRepository, validate *validator.Validate, kongClient *http.KongClient) *UserBaseUseCase {
	return &UserBaseUseCase{
		DB:             db,
		Log:            log,
		UserRepository: userRepo,
		Validate:       validate,
		KongClient:     kongClient,
	}
}

func (u *UserBaseUseCase) GenerateToken(user *model.User, secret string) (string, error) {
	claims := jwt.MapClaims{
		"iss":  user.ID,
		"data": user,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
