package model

import (
	"time"

	"github.com/google/uuid"
)

type Hotel struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"type:varchar;not null"`
	OwnerId     uuid.UUID `gorm:"type:uuid;not null"`
	Address     string    `gorm:"type:varchar;not null"`
	Description string    `gorm:"type:varchar;not null"`
	TotalRoom   uint64    `gorm:"type:bigint;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Hotel) TableName() string {
	return "hotels"
}
