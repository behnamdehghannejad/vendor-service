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

func (s *InventoryService) Upsert(inventoryRequest domain.Inventory) error {
	inventory, err := s.repository.GetInventory(inventoryRequest.VendorID, inventoryRequest.ProductID)
	if appErr, ok := err.(*apperror.AppError); ok && appErr.GetErrorType() != apperror.NotFound {
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
	if err := s.checkEnoughQuantityForReserving(
		reserveRequest.VendorID,
		reserveRequest.ProductID,
		reserveRequest.Reserved,
	); err != nil {
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

	err = iwf.IncreaseReserveInventory(
		reserveRequest.VendorID,
		reserveRequest.ProductID,
		reserveRequest.Reserved,
	)
	if err != nil {
		return err
	}

	err = iwf.CreateHistory(domain.History{
		ID:        reserveRequest.RequestID,
		Reserved:  reserveRequest.Reserved,
		VendorID:  reserveRequest.VendorID,
		ProductID: reserveRequest.ProductID,
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
