package model

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type HistoryModel struct {
	ID        int                  `gorm:"primary_key"`
	OrderID   string               `gorm:"column:order_id"`
	PaymentID string               `gorm:"column:payment_id"`
	Quantity  int                  `gorm:"column:quantity"`
	ProductID int                  `gorm:"column:product_id"`
	VendorID  int                  `gorm:"column:vendor_id"`
	Status    domain.HistoryStatus `gorm:"column:status"`
	Active    bool                 `gorm:"column:active"`
	CreatedAt time.Time            `gorm:column"created_at"`
	UpdatedAt time.Time            `gorm:column"updated_at"`
}

func (HistoryModel) TableName() string {
	return "histories" // or "history"
}
