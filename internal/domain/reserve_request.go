package domain

type ReserveRequest struct {
	VendorID  int
	ProductID int
	Reserved  int
	RequestID string
}
