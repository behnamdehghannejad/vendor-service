package port

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

type DiscountClient interface {
	GetDiscountPercentageProducts([]int) []domain.ProductDiscountPercentage
}
