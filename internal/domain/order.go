package domain

type ListOrder struct {
	Orders []Order
}

type Order struct {
	OrderID   string
	Quantity  int
	ProductID int
	VendorID  int
	PaymentID string
}
