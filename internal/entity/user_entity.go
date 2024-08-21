package entity

import (
	"time"

	"github.com/google/uuid"
)

// User is a struct that represents a user entity
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Token     string    `gorm:"column:token"`
	Email     string    `gorm:"column:email"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamptz;default:CURRENT_TIMESTAMP"`
}

func (u *User) TableName() string {
	return "users"
}
