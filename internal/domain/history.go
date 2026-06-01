package domain

import (
	"time"
)

type History struct {
	ID        int
	OrderID   string
	PaymentID string
	Quantity  int
	ProductID int
	VendorID  int
	Status    HistoryStatus
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HistoryStatus string

const (
	CREATED   HistoryStatus = "CREATED"
	RUNNING   HistoryStatus = "RUNNING"
	PAID      HistoryStatus = "PAID"
	READY     HistoryStatus = "READY"
	SENT      HistoryStatus = "SENT"
	DELIVERED HistoryStatus = "DELIVERED"
)
