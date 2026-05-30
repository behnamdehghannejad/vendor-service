package dto

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

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

type CreateHistoryRequest struct {
	OrderID   string `json:"order_id"`
	PaymentID string `json:"payment_id"`
	Quantity  int    `json:"quantity"`
	ProductID int    `json:"product_id"`
	VendorID  int    `json:"vendor_id"`
}

type SearchHistory struct {
	Activation string                `form:"activation"`
	PaymentID  string                `form:"payment_id"`
	OrderID    string                `form:"order_id"`
	VendorID   *int                  `form:"vendor_id" binding:"omitempty,gte=0"`
	ProductID  *int                  `form:"product_id" binding:"omitempty,gte=0"`
	Status     *domain.HistoryStatus `form:"status" binding:"omitempty,oneof=CREATED RUNNING PAID READY SENT DELIVERED"`
}

type ResponseHistories struct {
	Items []HistoryResponse `json:"items"`
}
