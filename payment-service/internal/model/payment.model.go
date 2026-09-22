package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentStatus string

const (
	PENDING PaymentStatus = "PENDING"
	SUCCESS PaymentStatus = "SUCCESS"
	FAIL    PaymentStatus = "FAIL"
)

type Payment struct {
	ID            uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid()"`
	BookingId     uuid.UUID       `gorm:"type:uuid;not null"`
	PaymentMethod string          `gorm:"type:varchar;not null"`
	Status        PaymentStatus   `gorm:"type:payment_status;not null"`
	TotalAmount   decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Payment) TableName() string {
	return "payments"
}
