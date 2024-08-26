package usecase

import (
	"context"
	"edot-monorepo/services/user-service/internal/entity"
	"edot-monorepo/services/user-service/internal/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterUseCase struct {
	*UserBaseUseCase
}

func NewUserRegisterUseCase(userBaseUseCase *UserBaseUseCase) *UserRegisterUseCase {
	return &UserRegisterUseCase{
		userBaseUseCase,
	}
}

func (c *UserRegisterUseCase) Exec(ctx context.Context, request *model.UserRegisterRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	total, err := c.UserRepository.CountByEmail(tx, request.Email)
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if total > 0 {
		c.Log.Warnf("User already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		Name:     request.Name,
		Password: string(password),
		Email:    request.Email,
	}

	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	token, err := c.GenerateToken(&model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, "")
	if err != nil {
		c.Log.Warnf("Error generate token : %+v", err)
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
