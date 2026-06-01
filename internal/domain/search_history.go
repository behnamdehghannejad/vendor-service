package domain

type SearchHistory struct {
	Activation *bool
	PaymentID  string
	OrderID    string
	VendorID   *int
	ProductID  *int
	Status     *HistoryStatus
}
