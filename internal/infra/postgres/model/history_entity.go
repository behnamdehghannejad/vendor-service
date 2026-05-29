package model

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/google/uuid"
)

type HistoryEntity struct {
	ID        int           `gorm:"primary_key"`
	OrderID   uuid.UUID     `gorm:"column:order_id"`
	PaymentID uuid.UUID     `gorm:"column:payment_id"`
	Quantity  int           `gorm:"column:quantity"`
	ProductID int           `gorm:"column:product_id"`
	VendorID  int           `gorm:"column:vendor_id"`
	Status    domain.Status `gorm:"column:status"`
	Active    bool          `gorm:"column:active"`
	CreatedAt time.Time     `gorm:column"created_at"`
	UpdatedAt time.Time     `gorm:column"updated_at"`
}

func (HistoryEntity) HistoryTableName() string {
	return "history"
}
