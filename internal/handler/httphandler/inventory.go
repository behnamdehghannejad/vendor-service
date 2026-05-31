package httphandler

import (
	"net/http"
	"strconv"

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

func NewInventoryHandler(inventoryService port.InventoryService, validator *validator.Inventory) *Inventory {
	return &Inventory{
		inventoryService: inventoryService,
		validator:        validator,
	}
}

func (i *Inventory) GetInventory(c *gin.Context) {
	ID, err := i.getIDs(c)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	inventory, err := i.inventoryService.FindInventory(ID)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, dto.ResponseInventory{
		Quantity:  inventory.Quantity,
		Reserved:  inventory.Reserved,
		VendorID:  inventory.VendorID,
		ProductID: inventory.ProductID,
	})
}

func (i *Inventory) getIDs(c *gin.Context) (int, error) {
	IDStr := c.Param("id")

	err := i.validator.ValidateIDs(IDStr)
	if err != nil {
		return 0, err
	}

	vendorID, _ := strconv.Atoi(IDStr)

	return vendorID, nil
}
