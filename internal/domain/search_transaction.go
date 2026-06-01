package domain

type SearchTransaction struct {
	Activation *bool
	PaymentID  string
	OrderID    string
	VendorID   *int
	ProductID  *int
	Status     *HistoryStatus
}
