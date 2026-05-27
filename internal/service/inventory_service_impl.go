package service

import (
	"vendor-service/internal/domain"
	"vendor-service/internal/infra/repository"
)

type InventoryServiceImpl struct {
	repository repository.InventoryRepository
}

func NewInventoryService(repository repository.InventoryRepository) *InventoryServiceImpl {
	return &InventoryServiceImpl{repository: repository}
}

func (service *InventoryServiceImpl) AddProductsToVendor(inventory *domain.Inventory) error {
	loadedInventory, err := service.FindByVendorIDAndProductID(inventory.VendorID, inventory.ProductID)
	if err != nil {
		inventory.Reserved = 0
		if err := service.repository.Add(inventory); err != nil {
			return err
		}
	}

	if loadedInventory != nil {
		loadedInventory.Quantity += inventory.Quantity
		if err := service.repository.Update(loadedInventory); err != nil {
			return err
		}
	}

	return nil
}

func (service *InventoryServiceImpl) FindByVendorIDAndProductID(vendorID int, productID int) (*domain.Inventory, error) {
	return service.repository.FindByVendorIDAndProductID(vendorID, productID)
}

func (service *InventoryServiceImpl) Update(inventory *domain.Inventory) error {
	return service.repository.Update(inventory)
}
