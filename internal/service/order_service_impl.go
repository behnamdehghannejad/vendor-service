package service

import (
	"vendor-service/internal/domain"
)

type OrderServiceImpl struct {
	inventoryService InventoryService
	productService   ProductService
	vendorService    VendorService
	historyService   HistoryService
}

func NewOrderServiceImpl(
	inventoryService InventoryService,
	productService ProductService,
	vendorService VendorService,
	historyService HistoryService,
) *OrderServiceImpl {
	return &OrderServiceImpl{
		inventoryService: inventoryService,
		productService:   productService,
		vendorService:    vendorService,
		historyService:   historyService,
	}
}

func (service *OrderServiceImpl) ManageOrders(orders *domain.ListOrder) error {
	for _, order := range orders.Orders {
		err := service.handleOrder(order)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *OrderServiceImpl) handleOrder(order domain.Order) error {
	if err := service.vendorService.IsActive(order.VendorID); err != nil {
		return err
	}

	if err := service.productService.IsActive(order.VendorID); err != nil {
		return err
	}

	err2 := service.findAndUpdateInventory(order)
	if err2 != nil {
		return err2
	}

	if err := service.historyService.Create(CreateHistoryByOrder(order)); err != nil {
		return err
	}

	return nil
}

func (service *OrderServiceImpl) findAndUpdateInventory(order domain.Order) error {
	inventory, err := service.inventoryService.FindByVendorIDAndProductID(order.VendorID, order.ProductID)
	if err != nil {
		return err
	}

	if err := domain.ValidateAndSetQuantity(inventory, order.Quantity); err != nil {
		return err
	}

	inventory.Reserved += order.Quantity
	if err := service.inventoryService.Update(inventory); err != nil {
		return err
	}
	return nil
}

func (service *OrderServiceImpl) AcceptOrdersPayment(orders *domain.ListOrder) error {
	for _, order := range orders.Orders {
		if err := service.updateInventory(order); err != nil {
			return err
		}

		if err := service.updateHistory(order); err != nil {
			return err
		}
	}

	return nil
}

func (service *OrderServiceImpl) updateInventory(order domain.Order) error {
	inventory, err := service.inventoryService.FindByOrderID(order.OrderID)
	if err != nil {
		return err
	}

	inventory.Reserved -= order.Quantity
	if err := service.inventoryService.Update(inventory); err != nil {
		return err
	}
	return nil
}

func (service *OrderServiceImpl) updateHistory(order domain.Order) error {
	history, err := service.historyService.FindByOrderID(order.OrderID)
	if err != nil {
		return err
	}

	history.Status = domain.PAID
	history.PaymentID = order.PaymentID
	if err := service.historyService.Update(history); err != nil {
		return err
	}
	return nil
}
