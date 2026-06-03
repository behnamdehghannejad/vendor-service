package domain

type Inventory struct {
	VendorID           int
	ProductID          int
	DiscountPercentage float64
	Quantity           int
	Reserved           int
	V                  int
	Price              int
}
