package httphandler

import (
	"net/http"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/httperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/gin-gonic/gin"
)

type History struct {
	service port.HistoryService
}

func NewHistoryHandler(service port.HistoryService) *History {
	return &History{
		service: service,
	}
}

func (h *History) Search(c *gin.Context) {
	var q dto.SearchHistory

	if err := c.ShouldBindQuery(&q); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	histories, err := h.service.Search(domain.SearchHistory{
		Activation: h.GetIsActiveFromQuery(q.Activation),
		PaymentID:  q.PaymentID,
		OrderID:    q.OrderID,
		VendorID:   q.VendorID,
		ProductID:  q.ProductID,
		Status:     q.Status,
	})
	if err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, dto.ResponseHistories{
		Items: h.serializeHistories(histories),
	})
}

func (*History) GetIsActiveFromQuery(activeStr string) *bool {
	active := true
	deActive := false
	switch activeStr {
	case "active":
		return &active
	case "deactive":
		return &deActive
	}
	return nil
}

func (h *History) serializeHistory(history domain.History) dto.HistoryResponse {
	return dto.HistoryResponse{
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
