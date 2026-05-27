package handler

import (
	"context"

	"vendor-service/internal/domain"
	"vendor-service/internal/service"
	pb "vendor-service/proto/generate"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderGrpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service service.OrderService
}

func NewOrderGrpcHandler(
	service service.OrderService,
) *OrderGrpcHandler {
	return &OrderGrpcHandler{
		service: service,
	}
}

func (handler *OrderGrpcHandler) AddOrders(ctx context.Context, request *pb.ManageOrdersRequest) (*emptypb.Empty, error) {
	requestDto := toManageOrdersDomain(request)

	if err := handler.service.ManageOrders(requestDto); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (handler *OrderGrpcHandler) AcceptOrdersPayment(ctx context.Context, request *pb.AcceptOrdersPaymentRequest) (*emptypb.Empty, error) {
	requestDto := toAcceptOrdersPaymentDomain(request)

	if err := handler.service.AcceptOrdersPayment(requestDto); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func toManageOrdersDomain(
	request *pb.ManageOrdersRequest,
) *domain.ListOrder {

	orders := make([]domain.Order, 0, len(request.Orders))

	for _, order := range request.Orders {
		orderID, _ := uuid.Parse(order.OrderId)
		domainOrder := domain.Order{
			OrderID:   orderID,
			ProductID: int(order.ProductId),
			VendorID:  int(order.VendorId),
			Quantity:  int(order.Quantity),
		}
		orders = append(orders, domainOrder)
	}

	return &domain.ListOrder{
		Orders: orders,
	}
}

func toAcceptOrdersPaymentDomain(request *pb.AcceptOrdersPaymentRequest) *domain.ListOrder {
	orders := make([]domain.Order, 0, len(request.Orders))

	for _, order := range request.Orders {
		orderID, _ := uuid.Parse(order.OrderId)
		paymentID, _ := uuid.Parse(order.PaymentId)

		domainOrder := domain.Order{
			OrderID:   orderID,
			PaymentID: paymentID,
			ProductID: int(order.ProductId),
			VendorID:  int(order.VendorId),
			Quantity:  int(order.Quantity),
		}

		orders = append(orders, domainOrder)
	}

	return &domain.ListOrder{Orders: orders}
}
