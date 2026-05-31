package dto

type ResponseInventory struct {
	Reserved  int `json:"reserved"`
	VendorID  int `json:"vendor_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
