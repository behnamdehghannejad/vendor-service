package domain

import (
	"time"
)

type History struct {
	ID        string
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
	DRAFTED HistoryStatus = "DRAFTED"
	SUCCESS HistoryStatus = "SUCCESS"
)
