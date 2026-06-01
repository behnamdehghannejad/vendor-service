package dto

type OrderRequest struct {
	OrderID   string `json:"order_id"`
	PaymentID string `json:"payment_id"`
	VendorID  int    `json:"vendor_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type ManageOrdersRequest struct {
	Orders []OrderRequest `json:"orders"`
}

type AcceptOrdersPaymentRequest struct {
	Orders []OrderRequest `json:"orders"`
}
