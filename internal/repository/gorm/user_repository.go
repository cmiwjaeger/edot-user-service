package repository

import (
	"edot-monorepo/services/user-service/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (repo *UserRepository) FindByEmail(tx *gorm.DB, user *entity.User, email string) error {
	return tx.Where("email = ? ", email).First(user).Error
}

func (r *UserRepository) CountByEmail(db *gorm.DB, email string) (int64, error) {
	var total int64
	err := db.Model(new(entity.User)).Where("email = ?", email).Count(&total).Error
	return total, err
}
