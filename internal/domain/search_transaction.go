package domain

type SearchTransaction struct {
	Activation *bool
	VendorID   *int
	ProductID  *int
	Status     *TransactionStatus
}
