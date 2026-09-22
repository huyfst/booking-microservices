package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	UserName  string    `gorm:"type:varchar;not null"`
	Email     string    `gorm:"type:varchar;not null"`
	Password  string    `gorm:"type:varchar;not null"`
	CreatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
