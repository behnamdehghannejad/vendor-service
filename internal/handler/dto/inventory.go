package dto

type ResponseInventories struct {
	Items []ResponseInventory `json:"items"`
}

type ResponseInventory struct {
	Reserved  int `json:"reserved"`
	VendorID  int `json:"vendor_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type RequestUpsertInventory struct {
	Quantity int `json:"quantity"`
}

type RequestReserve struct {
	Reserve   int    `json:"reserve"`
	RequestID string `json:"request_id"`
}

type SearchInventory struct {
	VendorID  *int `form:"vendor_id"`
	ProductID *int `form:"product_id"`
}
