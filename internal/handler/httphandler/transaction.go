package httphandler

import (
	"net/http"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/handler/dto"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/httperror"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	service port.TransactionService
}

func NewTransactionHandler(service port.TransactionService) *Transaction {
	return &Transaction{
		service: service,
	}
}

func (h *Transaction) Search(c *gin.Context) {
	var q dto.SearchTransaction

	if err := c.ShouldBindQuery(&q); err != nil {
		errorResponse, status := httperror.Handle(err)
		c.JSON(status, errorResponse)
		return
	}

	transactions, err := h.service.Search(domain.SearchTransaction{
		Activation: h.GetIsActiveFromQuery(q.Activation),
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
		Items: h.serializeHistories(transactions),
	})
}

func (*Transaction) GetIsActiveFromQuery(activeStr string) *bool {
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

func (h *Transaction) serializeTransaction(transaction domain.Transaction) dto.TransactionResponse {
	return dto.TransactionResponse{
		Reserved:  transaction.Reserved,
		ProductID: transaction.ProductID,
		VendorID:  transaction.VendorID,
		Status:    string(transaction.Status),
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
	}
}

func (h *Transaction) serializeHistories(transactions []domain.Transaction) []dto.TransactionResponse {
	items := make([]dto.TransactionResponse, 0, len(transactions))

	for _, transaction := range transactions {
		items = append(items, h.serializeTransaction(transaction))
	}

	return items
}
