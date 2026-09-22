package model

import (
	"time"

	"github.com/google/uuid"
)

type Inventory struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	HotelId       uuid.UUID `gorm:"type:uuid;not null"`
	AvailableRoom uint64    `gorm:"type:bigint;not null"`
	Date          time.Time `gorm:"type:timestamp with time zone;not null;unique"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Inventory) TableName() string {
	return "inventories"
}
