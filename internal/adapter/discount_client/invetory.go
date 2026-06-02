package discount

import "github.com/behnamdehghannejad/vendorservice/internal/domain"

func (d *DiscountClient) GetDiscountPercentageProducts(inventoryIdentifies []domain.InventoryIdentity) []domain.InventoryDiscountPercentage {
	productDiscountPercentages := make([]domain.InventoryDiscountPercentage, 0, len(inventoryIdentifies))
	for _, inventoryIdentity := range inventoryIdentifies {
		productDiscountPercentages = append(productDiscountPercentages, domain.InventoryDiscountPercentage{
			ProductID:          inventoryIdentity.ProductID,
			VendorID:           inventoryIdentity.VendorID,
			DiscountPercentage: 0,
		})
	}
	return productDiscountPercentages
}
