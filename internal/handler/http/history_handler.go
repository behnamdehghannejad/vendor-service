package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"vendor-service/internal/domain"
	"vendor-service/internal/handler/dto"
	"vendor-service/internal/service"

	"github.com/google/uuid"
)

type HistoryHandler struct {
	service service.HistoryService
}

func NewHistoryHandler(service service.HistoryService) *HistoryHandler {
	return &HistoryHandler{service: service}
}

func (handler *HistoryHandler) Create(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var req dto.CreateHistoryRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.service.Create(toHistoryDomain(req))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer)
}

func (handler *HistoryHandler) GetByOrderID(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	id := request.PathValue("id")
	uid, err := toUUID(writer, id)

	history, err := handler.service.FindByOrderID(uid)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	returnResponse(writer, history)
}

func (handler *HistoryHandler) GetByPaymentID(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	id := request.PathValue("id")
	uid, err := toUUID(writer, id)

	history, err := handler.service.FindByPaymentID(uid)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	returnResponse(writer, history)
}

func (handler *HistoryHandler) GetByProductID(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	histories, err := handler.service.FindByProductID(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(toHistoryResponseList(histories))
}

func (handler *HistoryHandler) GetByVendorID(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	histories, err := handler.service.FindByVendorID(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(toHistoryResponseList(histories))
}

func (handler *HistoryHandler) GetByStatus(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	status := request.PathValue("status")

	histories, err := handler.service.FindByStatus(domain.Status(status))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(toHistoryResponseList(histories))
}

func (handler *HistoryHandler) GetByIsActive(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	activeStr := request.PathValue("active")
	active, err := strconv.ParseBool(activeStr)
	if err != nil {
		http.Error(writer, "invalid active value", http.StatusBadRequest)
		return
	}

	histories, err := handler.service.FindByIsActive(active)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(toHistoryResponseList(histories))
}

func (handler *HistoryHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	if err := handler.service.Delete(id); err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
	json.NewEncoder(writer)
}

func returnResponse(writer http.ResponseWriter, history *domain.History) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(toHistoryResponse(&history))
}

func toHistoryDomain(req dto.CreateHistoryRequest) *domain.History {
	orderID, _ := uuid.Parse(req.OrderID)
	paymentID, _ := uuid.Parse(req.PaymentID)

	return &domain.History{
		OrderID:   orderID,
		PaymentID: paymentID,
		Quantity:  req.Quantity,
		ProductID: req.ProductID,
		VendorID:  req.VendorID,
		Status:    domain.CREATED,
		Active:    true,
	}
}

func toHistoryResponseList(all []domain.History) []dto.HistoryResponse {
	result := make([]dto.HistoryResponse, 0, len(all))

	for _, h := range all {
		result = append(result, toHistoryResponse(h))
	}

	return result
}

func toHistoryResponse(history *domain.History) dto.HistoryResponse {

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

func toUUID(writer http.ResponseWriter, id string) (uuid.UUID, error) {
	parse, err := uuid.Parse(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	return parse, err
}
