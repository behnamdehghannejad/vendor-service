package service

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
)

type InventoryService struct {
	repository port.InventoryRepository
}

func NewInventoryService(repository port.InventoryRepository) *InventoryService {
	return &InventoryService{repository: repository}
}

func (s *InventoryService) AddProductsToVendor(inventory domain.Inventory) error {
	loadedInventory, err := s.FindByVendorIDAndProductID(inventory.VendorID, inventory.ProductID)
	if err != nil {
		inventory.Reserved = 0
		if err := s.repository.Create(inventory); err != nil {
			return err
		}
	}

	if loadedInventory.ID != 0 {
		loadedInventory.Quantity += inventory.Quantity
		if err := s.repository.Update(loadedInventory); err != nil {
			return err
		}
	}

	return nil
}

func (s *InventoryService) FindByVendorIDAndProductID(vendorID int, productID int) (domain.Inventory, error) {
	return s.repository.FindByVendorIDAndProductID(vendorID, productID)
}

func (s *InventoryService) Update(inventory domain.Inventory) error {
	return s.repository.Update(inventory)
}
