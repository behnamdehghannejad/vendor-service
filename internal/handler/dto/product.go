package dto

import "time"

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductsResponse struct {
	Items []ProductResponse `json:"items"`
}

type RequestUpdateProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
