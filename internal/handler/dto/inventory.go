package dto

type AddProductsToVendorRequest struct {
	ProductID int `json:"product_id"`
	VendorID  int `json:"vendor_id"`
	Quantity  int `json:"quantity"`
}
