package handler

import (
	"context"
	"time"
	"vendor-service/internal/domain"
	"vendor-service/internal/service"
	pb "vendor-service/proto/generate"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HistoryGrpcHandler struct {
	pb.UnimplementedHistoryServiceServer
	service service.HistoryService
}

func NewHistoryGrpcHandler(service service.HistoryService) *HistoryGrpcHandler {
	return &HistoryGrpcHandler{
		service: service,
	}
}

func (handler *HistoryGrpcHandler) Add(ctx context.Context, request *pb.CreateHistoryRequest) (*emptypb.Empty, error) {

	requestDto := toHistoryDomain(request)

	if err := handler.service.Create(requestDto); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (handler *HistoryGrpcHandler) Update(ctx context.Context, request *pb.UpdateHistoryRequest) (*emptypb.Empty, error) {
	history := historyProtoToDomain(request.History)

	if err := handler.service.Update(&history); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (handler *HistoryGrpcHandler) Delete(ctx context.Context, request *pb.DeleteHistoryRequest) (*emptypb.Empty, error) {
	if err := handler.service.Delete(int(request.Id)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (handler *HistoryGrpcHandler) FindByOrderID(ctx context.Context, request *pb.GetHistoryByOrderIDRequest) (*pb.HistoryResponse, error) {

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	history, err := handler.service.FindByOrderID(id)
	if err != nil {
		return nil, err
	}

	return &pb.HistoryResponse{
		History: toHistoryProto(*history),
	}, nil
}

func (handler *HistoryGrpcHandler) FindByPaymentID(ctx context.Context, request *pb.GetHistoryByPaymentIDRequest) (*pb.HistoryResponse, error) {

	id, err := uuid.Parse(request.PaymentId)
	if err != nil {
		return nil, err
	}

	history, err := handler.service.FindByPaymentID(id)
	if err != nil {
		return nil, err
	}

	return &pb.HistoryResponse{
		History: toHistoryProto(*history),
	}, nil
}

func (handler *HistoryGrpcHandler) FindByProductID(ctx context.Context, request *pb.GetHistoryByProductIDRequest) (*pb.ListHistoryResponse, error) {

	histories, err := handler.service.FindByProductID(int(request.ProductId))
	if err != nil {
		return nil, err
	}

	var result []*pb.History
	for _, h := range histories {
		result = append(result, toHistoryProto(h))
	}

	return &pb.ListHistoryResponse{
		Histories: result,
	}, nil
}

func (handler *HistoryGrpcHandler) FindByVendorID(ctx context.Context, request *pb.GetHistoryByVendorIDRequest) (*pb.ListHistoryResponse, error) {

	histories, err := handler.service.FindByVendorID(int(request.VendorId))
	if err != nil {
		return nil, err
	}

	var result []*pb.History
	for _, h := range histories {
		result = append(result, toHistoryProto(h))
	}

	return &pb.ListHistoryResponse{
		Histories: result,
	}, nil
}

func (handler *HistoryGrpcHandler) FindByStatus(ctx context.Context, request *pb.GetHistoryByStatusRequest) (*pb.ListHistoryResponse, error) {

	histories, err := handler.service.FindByStatus(domain.Status(request.Status.String()))
	if err != nil {
		return nil, err
	}

	var result []*pb.History
	for _, h := range histories {
		result = append(result, toHistoryProto(h))
	}

	return &pb.ListHistoryResponse{
		Histories: result,
	}, nil
}

func (handler *HistoryGrpcHandler) FindByIsActive(ctx context.Context, request *pb.GetHistoryByActiveRequest) (*pb.ListHistoryResponse, error) {

	histories, err := handler.service.FindByIsActive(request.Active)
	if err != nil {
		return nil, err
	}

	var result []*pb.History
	for _, h := range histories {
		result = append(result, toHistoryProto(h))
	}

	return &pb.ListHistoryResponse{
		Histories: result,
	}, nil
}

func toHistoryDomain(request *pb.CreateHistoryRequest) *domain.History {

	orderID, _ := uuid.Parse(request.OrderId)
	paymentID, _ := uuid.Parse(request.PaymentId)

	return &domain.History{
		OrderID:   orderID,
		PaymentID: paymentID,
		Quantity:  int(request.Quantity),
		ProductID: int(request.ProductId),
		VendorID:  int(request.VendorId),
		Status:    domain.CREATED,
		Active:    true,
	}
}

func historyProtoToDomain(history *pb.History) domain.History {

	orderID, _ := uuid.Parse(history.OrderId)
	paymentID, _ := uuid.Parse(history.PaymentId)

	return domain.History{
		ID:        int(history.Id),
		OrderID:   orderID,
		PaymentID: paymentID,
		Quantity:  int(history.Quantity),
		ProductID: int(history.ProductId),
		VendorID:  int(history.VendorId),
		Status:    domain.Status(history.Status.String()),
		Active:    history.Active,
	}
}

func toHistoryProto(history domain.History) *pb.History {
	return &pb.History{
		Id:        int32(history.ID),
		OrderId:   history.OrderID.String(),
		PaymentId: history.PaymentID.String(),
		Quantity:  int32(history.Quantity),
		ProductId: int32(history.ProductID),
		VendorId:  int32(history.VendorID),
		Status:    pb.Status(pb.Status_value[string(history.Status)]),
		Active:    history.Active,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
	}
}
