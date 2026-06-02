package domain

import (
	"time"
)

type Transaction struct {
	ID        string
	Reserved  int
	ProductID int
	VendorID  int
	Status    TransactionStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionStatus string

const (
	HISTORY_DRAFT   TransactionStatus = "DRAFTED"
	HISTORY_SUCCESS TransactionStatus = "SUCCESS"
	HISTORY_FAIL    TransactionStatus = "FAIL"
)
