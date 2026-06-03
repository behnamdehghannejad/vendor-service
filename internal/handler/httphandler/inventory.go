package httphandler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/httperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/behnamdehghannejad/vendorservice/internal/validator"
	"github.com/gin-gonic/gin"
)

type Inventory struct {
	inventoryService port.InventoryService
	validator        *validator.Inventory
}

func NewInventory(inventoryService port.InventoryService, validator *validator.Inventory) *Inventory {
	return &Inventory{
		inventoryService: inventoryService,
		validator:        validator,
	}
}

func (i *Inventory) GetInventory(c *gin.Context) {
	vendorID, productID, err := i.GetIDs(c)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	inventory, err := i.inventoryService.FindInventory(vendorID, productID)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, i.serializeInventory(inventory))
}

func (i *Inventory) Upsert(c *gin.Context) {
	vendorID, productID, err := i.GetIDs(c)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	var req dto.RequestUpsertInventory
	if err := c.ShouldBind(&req); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	if err := i.validator.Upsert(req); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	err = i.inventoryService.Upsert(domain.Inventory{
		VendorID:  vendorID,
		ProductID: productID,
		Quantity:  req.Quantity,
		Price:     req.Price,
	})
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.Status(http.StatusNoContent)
}

func (i *Inventory) Search(c *gin.Context) {
	var req dto.SearchInventory
	if err := c.ShouldBindQuery(&req); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	inventories, err := i.inventoryService.Search(domain.SearchInventory{
		VendorID:  req.VendorID,
		ProductID: req.ProductID,
	})
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}
	c.JSON(http.StatusOK, dto.ResponseInventories{
		Items: i.serializeInventories(inventories),
	})
}

func (i *Inventory) Reserve(c *gin.Context) {
	vendorID, productID, err := i.GetIDs(c)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	var req dto.RequestReserve
	if err := c.ShouldBind(&req); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	err = i.inventoryService.ReserveQuantity(domain.ReserveRequest{
		VendorID:  vendorID,
		ProductID: productID,
		Reserved:  req.Reserve,
		RequestID: req.RequestID,
	})
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.Status(http.StatusNoContent)
}

func (i *Inventory) GetIDs(c *gin.Context) (int, int, error) {
	ids := strings.Split(c.Param("vpIDs"), "-")
	err := i.validator.ValidateIDs(ids[0], ids[1])
	if err != nil {
		return 0, 0, err
	}

	vendorID, _ := strconv.Atoi(ids[0])
	productID, _ := strconv.Atoi(ids[1])

	return vendorID, productID, nil
}

func (i *Inventory) serializeInventories(inventories []domain.Inventory) []dto.ResponseInventory {
	inventoriesResponse := make([]dto.ResponseInventory, 0, len(inventories))
	for _, inventory := range inventories {
		inventoriesResponse = append(inventoriesResponse, i.serializeInventory(inventory))
	}
	return inventoriesResponse
}

func (i *Inventory) serializeInventory(inventory domain.Inventory) dto.ResponseInventory {
	return dto.ResponseInventory{
		VendorID:           inventory.VendorID,
		ProductID:          inventory.ProductID,
		Quantity:           inventory.Quantity,
		DiscountPercentage: inventory.DiscountPercentage,
		Reserved:           inventory.Reserved,
		Price:              inventory.Price,
	}
}
