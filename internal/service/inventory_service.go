package service

import (
	"context"
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/log"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
)

type InventoryService struct {
	repository        port.InventoryRepository
	discountClient    port.DiscountClient
	unitOfWorkFactory port.UnitOfWorkFactor
}

func NewInventoryService(
	repository port.InventoryRepository,
	unitOfWorkFactor port.UnitOfWorkFactor,
	discountClient port.DiscountClient,
) *InventoryService {
	return &InventoryService{
		discountClient:    discountClient,
		repository:        repository,
		unitOfWorkFactory: unitOfWorkFactor,
	}
}

func (s *InventoryService) FindInventory(vendorID int, productID int) (domain.Inventory, error) {
	return s.repository.GetInventory(vendorID, productID)
}

func (s *InventoryService) Search(search domain.SearchInventory) ([]domain.Inventory, error) {
	return s.repository.Filter(search)
}

func (s *InventoryService) Upsert(inventoryRequest domain.Inventory) error {
	inventory, err := s.repository.GetInventory(inventoryRequest.VendorID, inventoryRequest.ProductID)
	if appErr, ok := err.(*apperror.AppError); err != nil && (!ok || appErr.GetErrorType() != apperror.NotFound) {
		return err
	}

	if inventory.Reserved > inventoryRequest.Quantity {
		return apperror.Wrap(err).
			BadRequest().
			Warningf("the request quantity must bel less than reserved").
			Build()
	}

	inventoryRequest.V = inventory.V

	return s.repository.Upsert(inventoryRequest)
}

func (s *InventoryService) ReserveQuantity(reserveRequest domain.ReserveRequest) error {
	inventory, err := s.repository.GetInventory(reserveRequest.VendorID, reserveRequest.ProductID)
	if err != nil {
		return err
	}

	if inventory.Reserved+reserveRequest.Reserved > inventory.Quantity {
		return apperror.WithoutParentError().
			BadRequest().
			Warningf("the quantity isn't adequate").
			Build()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	iwf, err := s.unitOfWorkFactory.CreateInventoryUnitOfWork(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			return
		}

		if rollbackErr := iwf.Rollback(); rollbackErr != nil {
			log.Warningf("rollback failed: %v", rollbackErr)
		}
	}()

	err = iwf.IncreaseReserveInventory(
		domain.RequestReserve{
			VendorID:  reserveRequest.VendorID,
			ProductID: reserveRequest.ProductID,
			Reserved:  reserveRequest.Reserved,
			V:         inventory.V,
		},
	)
	if err != nil {
		return err
	}

	err = iwf.CreateTransaction(domain.Transaction{
		ID:        reserveRequest.RequestID,
		Reserved:  reserveRequest.Reserved,
		VendorID:  reserveRequest.VendorID,
		ProductID: reserveRequest.ProductID,
		Status:    domain.TRANSACTION_DRAFT,
	})
	if err != nil {
		return err
	}

	err = iwf.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *InventoryService) UpdateAllInventoriesDiscountPercentage() {
	inventories, err := s.repository.Filter(domain.SearchInventory{})
	if err != nil {
		return
	}

	productDiscountPercentages := s.discountClient.GetDiscountPercentageProducts(
		s.getInventoryKeys(inventories),
	)

	s.repository.UpdateProductDiscountPercentages(productDiscountPercentages)
}

func (*InventoryService) getInventoryKeys(inventories []domain.Inventory) []domain.InventoryIdentity {
	inventoryKeys := make([]domain.InventoryIdentity, 0, len(inventories))
	for _, inventory := range inventories {
		inventoryKeys = append(inventoryKeys, domain.InventoryIdentity{
			ProductID: inventory.ProductID,
			VendorID:  inventory.VendorID,
		})
	}
	return inventoryKeys
}
