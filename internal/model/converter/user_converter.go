package converter

import (
	"edot-monorepo/services/user-service/internal/entity"
	"edot-monorepo/services/user-service/internal/model"
)

func UserToResponse(user *entity.User, token string) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToTokenResponse(user *entity.User, token string) *model.UserResponse {
	return &model.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}
}
