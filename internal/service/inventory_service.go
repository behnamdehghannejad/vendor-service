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
	unitOfWorkFactory port.UnitOfWorkFactor
}

func NewInventoryService(repository port.InventoryRepository, unitOfWorkFactor port.UnitOfWorkFactor) *InventoryService {
	return &InventoryService{
		repository:        repository,
		unitOfWorkFactory: unitOfWorkFactor,
	}
}

func (s *InventoryService) FindInventory(vendorID int, productID int) (domain.Inventory, error) {
	return s.repository.GetInventory(vendorID, productID)
}

func (s *InventoryService) Upsert(inventory domain.Inventory) error {
	return s.repository.Upsert(inventory)
}

func (s *InventoryService) ReserveQuantity(vendorID int, productID int, reserved int, requestID string) error {
	if err := s.checkEnoughQuantityForReserving(vendorID, productID, reserved); err != nil {
		return err
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

	err = iwf.IncreaseReserveInventory(vendorID, productID, reserved)
	if err != nil {
		return err
	}

	err = iwf.CreateHistory(domain.History{
		ID:        requestID,
		Reserved:  reserved,
		VendorID:  vendorID,
		ProductID: productID,
		Status:    domain.HISTORY_DRAFT,
	})
	if err != nil {
		return err
	}

	return iwf.Commit()
}

func (s *InventoryService) checkEnoughQuantityForReserving(vendorID int, productID int, reserved int) error {
	inventory, err := s.repository.GetInventory(vendorID, productID)
	if err != nil {
		return err
	}

	if inventory.Reserved+reserved > inventory.Quantity {
		return apperror.WithoutParentError().
			BadRequest().
			Warningf("the quantity isn't adequate").
			Build()
	}
	return err
}
