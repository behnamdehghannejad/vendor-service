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
	TRANSACTION_DRAFT   TransactionStatus = "DRAFTED"
	TRANSACTION_SUCCESS TransactionStatus = "SUCCESS"
	TRANSACTION_FAIL    TransactionStatus = "FAIL"
)
