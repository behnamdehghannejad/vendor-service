package domain

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	ID        int
	OrderID   uuid.UUID
	PaymentID uuid.UUID
	Quantity  int
	ProductID int
	VendorID  int
	Status    Status
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status string

const (
	CREATED   = "CREATED"
	RUNNING   = "RUNNING"
	PAID      = "PAID"
	READY     = "READY"
	SENT      = "SENT"
	DELIVERED = "DELIVERED"
)
