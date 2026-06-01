package model

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type HistoryModel struct {
	ID        string               `gorm:"primaryKey;column:id"`
	Reserved  int                  `gorm:"column:reserved"`
	ProductID int                  `gorm:"column:product_id"`
	VendorID  int                  `gorm:"column:vendor_id"`
	Status    domain.HistoryStatus `gorm:"column:status"`
	CreatedAt time.Time            `gorm:"column:created_at"`
	UpdatedAt time.Time            `gorm:"column:updated_at"`
}

func (HistoryModel) TableName() string {
	return "histories"
}
