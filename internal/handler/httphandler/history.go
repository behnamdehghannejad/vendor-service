package httphandler

import (
	"net/http"
	"strconv"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/httperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/behnamdehghannejad/vendorservice/internal/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type History struct {
	service   port.HistoryService
	validator *validator.History
}

func NewHistoryHandler(service port.HistoryService, validator *validator.History) *History {
	return &History{
		service:   service,
		validator: validator,
	}
}

func (h *History) GetByOrderID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	history, err := h.service.FindByOrderID(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, h.serializeHistory(history))
}

func (h *History) GetByPaymentID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	history, err := h.service.FindByPaymentID(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, h.serializeHistory(history))
}

func (h *History) GetByProductID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	histories, err := h.service.FindByProductID(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, h.serializeHistories(histories))
}

func (h *History) GetByVendorID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	histories, err := h.service.FindByVendorID(id)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, h.serializeHistories(histories))
}

func (h *History) GetByStatus(c *gin.Context) {
	status := c.Param("status")

	histories, err := h.service.FindByStatus(domain.Status(status))
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, h.serializeHistories(histories))
}

func (h *History) GetByIsActive(c *gin.Context) {
	activeStr := c.Param("active")

	active, err := strconv.ParseBool(activeStr)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	histories, err := h.service.FindByIsActive(active)
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, h.serializeHistories(histories))
}

func (h *History) toHistoryDomain(req dto.CreateHistoryRequest) domain.History {
	orderID, _ := uuid.Parse(req.OrderID)
	paymentID, _ := uuid.Parse(req.PaymentID)

	return domain.History{
		OrderID:   orderID,
		PaymentID: paymentID,
		Quantity:  req.Quantity,
		ProductID: req.ProductID,
		VendorID:  req.VendorID,
		Status:    domain.CREATED,
		Active:    true,
	}
}

func (h *History) serializeHistory(history domain.History) dto.HistoryResponse {
	return dto.HistoryResponse{
		OrderID:   history.OrderID.String(),
		PaymentID: history.PaymentID.String(),
		Quantity:  history.Quantity,
		ProductID: history.ProductID,
		VendorID:  history.VendorID,
		Status:    string(history.Status),
		Active:    history.Active,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
	}
}

func (h *History) serializeHistories(histories []domain.History) []dto.HistoryResponse {
	items := make([]dto.HistoryResponse, 0, len(histories))

	for _, history := range histories {
		items = append(items, h.serializeHistory(history))
	}

	return items
}
