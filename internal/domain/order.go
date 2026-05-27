package domain

import (
	"github.com/google/uuid"
)

type ListOrder struct {
	Orders []Order
}

type Order struct {
	OrderID   uuid.UUID
	Quantity  int
	ProductID int
	VendorID  int
	PaymentID uuid.UUID
}
