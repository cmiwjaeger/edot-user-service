package usecase

import (
	"context"
	"edot-monorepo/services/user-service/internal/entity"
	"edot-monorepo/services/user-service/internal/model"
	"edot-monorepo/services/user-service/internal/model/converter"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserLoginUseCase struct {
	*UserBaseUseCase
}

func NewUserLoginUseCase(userBaseUseCase *UserBaseUseCase) *UserLoginUseCase {
	return &UserLoginUseCase{
		userBaseUseCase,
	}
}

func (c *UserLoginUseCase) Exec(ctx context.Context, request *model.UserLoginRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body  : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := c.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.Log.Warnf("Failed to compare user password with bcrype hash : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	user.Token = uuid.New().String()
	if err := c.UserRepository.Update(tx, user); err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToTokenResponse(user), nil
}

func generateJWT(userID string, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
