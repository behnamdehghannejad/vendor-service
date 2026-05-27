package service

import (
	"vendor-service/internal/domain"
	"vendor-service/internal/infra/repository"

	"github.com/google/uuid"
)

type InventoryServiceImpl struct {
	repository repository.InventoryRepository
}

func NewInventoryServiceImpl(repository repository.InventoryRepository) *InventoryServiceImpl {
	return &InventoryServiceImpl{repository: repository}
}

func (service *InventoryServiceImpl) FindByVendorIDAndProductID(vendorID int, productID int) (*domain.Inventory, error) {
	return service.repository.FindByVendorIDAndProductID(vendorID, productID)
}

func (service *InventoryServiceImpl) Update(inventory *domain.Inventory) error {
	return service.repository.Update(inventory)
}

func (service *InventoryServiceImpl) FindByOrderID(orderID uuid.UUID) (*domain.Inventory, error) {
	return service.repository.FindByOrderID(orderID)
}
