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

	c.JSON(http.StatusOK, dto.ResponseInventory{
		Quantity:  inventory.Quantity,
		Reserved:  inventory.Reserved,
		VendorID:  inventory.VendorID,
		ProductID: inventory.ProductID,
	})
}

func (i *Inventory) GetIDs(c *gin.Context) (int, int, error) {
	vendorIDStr := c.Param("vendor_id")
	productIDStr := c.Param("product_id")

	err := i.validator.ValidateIDs(vendorIDStr, productIDStr)
	if err != nil {
		return 0, 0, err
	}

	vendorID, _ := strconv.Atoi(vendorIDStr)
	productID, _ := strconv.Atoi(productIDStr)

	return vendorID, productID, nil
}
