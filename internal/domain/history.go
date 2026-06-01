package domain

import (
	"time"
)

type History struct {
	ID        string
	Reserved  int
	ProductID int
	VendorID  int
	Status    HistoryStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HistoryStatus string

const (
	HISTORY_DRAFT   HistoryStatus = "DRAFTED"
	HISTORY_SUCCESS HistoryStatus = "SUCCESS"
	HISTORY_FAIL    HistoryStatus = "FAIL"
)
