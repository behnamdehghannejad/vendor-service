package discount

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

func (d *DiscountClient) GetDiscountPercentageProducts(IDs []int) []domain.ProductDiscountPercentage {
	productDiscountPercentages := make([]domain.ProductDiscountPercentage, 0, len(IDs))
	for _, ID := range IDs {
		productDiscountPercentages = append(productDiscountPercentages, domain.ProductDiscountPercentage{
			ProductID:          ID,
			DiscountPercentage: 0,
		})
	}
	return productDiscountPercentages
}
