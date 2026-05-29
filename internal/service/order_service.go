package service

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

type OrderService struct {
	inventoryService InventoryService
	productService   ProductService
	vendorService    VendorService
	historyService   HistoryService
}

func NewOrderService(
	inventoryService InventoryService,
	productService ProductService,
	vendorService VendorService,
	historyService HistoryService,
) *OrderService {
	return &OrderService{
		inventoryService: inventoryService,
		productService:   productService,
		vendorService:    vendorService,
		historyService:   historyService,
	}
}

func (service *OrderService) ManageOrders(orders domain.ListOrder) error {
	for _, order := range orders.Orders {
		err := service.handleOrder(order)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *OrderService) handleOrder(order domain.Order) error {
	if err := service.vendorService.IsActive(order.VendorID); err != nil {
		return err
	}

	if err := service.productService.IsActive(order.VendorID); err != nil {
		return err
	}

	if err := service.findAndUpdateInventory(order); err != nil {
		return err
	}

	if err := service.historyService.Create(CreateHistoryByOrder(order)); err != nil {
		return err
	}

	return nil
}

func (service *OrderService) findAndUpdateInventory(order domain.Order) error {
	inventory, err := service.inventoryService.FindByVendorIDAndProductID(order.VendorID, order.ProductID)
	if err != nil {
		return err
	}

	if err := domain.ValidateAndSetQuantity(&inventory, order.Quantity); err != nil {
		return err
	}

	inventory.Reserved += order.Quantity
	if err := service.inventoryService.Update(inventory); err != nil {
		return err
	}
	return nil
}

func (service *OrderService) AcceptOrdersPayment(orders domain.ListOrder) error {
	for _, order := range orders.Orders {
		history, err := service.historyService.FindByOrderID(order.OrderID)
		if err != nil {
			return err
		}

		history.Status = domain.PAID
		history.PaymentID = order.PaymentID
		if err := service.historyService.Update(history); err != nil {
			return err
		}

		inventory, err := service.inventoryService.FindByVendorIDAndProductID(history.VendorID, history.ProductID)
		if err != nil {
			return err
		}

		inventory.Reserved -= history.Quantity
		if err := service.inventoryService.Update(inventory); err != nil {
			return err
		}
	}

	return nil
}

func CreateHistoryByOrder(order domain.Order) domain.History {
	return domain.History{
		OrderID:   order.OrderID,
		Quantity:  order.Quantity,
		ProductID: order.ProductID,
		VendorID:  order.VendorID,
		Status:    domain.CREATED,
		Active:    true,
		CreatedAt: time.Now(),
	}
}
