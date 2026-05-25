package dto

import "time"

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateProductRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateVendorRequest struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UpdateVendorRequest struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Active  bool   `json:"active"`
}

type VendorResponse struct {
	ID        int       `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateHistoryRequest struct {
	OrderID   string `json:"order_id"`
	PaymentID string `json:"payment_id"`
	Quantity  int    `json:"quantity"`
	ProductID int    `json:"product_id"`
	VendorID  int    `json:"vendor_id"`
}

type HistoryResponse struct {
	OrderID   string    `json:"order_id"`
	PaymentID string    `json:"payment_id"`
	Quantity  int       `json:"quantity"`
	ProductID int       `json:"product_id"`
	VendorID  int       `json:"vendor_id"`
	Status    string    `json:"status"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HistoryStatusRequest struct {
	Status string `json:"status"`
}

type HistoryActiveRequest struct {
	Active bool `json:"active"`
}
