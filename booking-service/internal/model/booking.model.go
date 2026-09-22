package model

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	UserId      uuid.UUID `gorm:"type:uuid;not null"`
	BookingCode string    `gorm:"type:varchar;not null"`
	HotelId     uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Booking) TableName() string {
	return "bookings"
}
