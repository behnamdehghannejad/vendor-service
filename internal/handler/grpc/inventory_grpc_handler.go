package handler

import (
	"context"
	"vendor-service/internal/domain"
	"vendor-service/internal/service"
	pb "vendor-service/proto/generate"

	"google.golang.org/protobuf/types/known/emptypb"
)

type InventoryGrpcHandler struct {
	pb.UnimplementedInventoryServiceServer
	service service.InventoryService
}

func NewInventoryGrpcHandler(service service.InventoryService) *InventoryGrpcHandler {
	return &InventoryGrpcHandler{
		service: service,
	}
}

func (handler *InventoryGrpcHandler) AddProductsToVendor(ctx context.Context, request *pb.AddProductsToVendorRequest) (*emptypb.Empty, error) {
	requestDto := toInventoryDomain(request)
	if err := handler.service.AddProductsToVendor(requestDto); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func toInventoryDomain(request *pb.AddProductsToVendorRequest) *domain.Inventory {
	return &domain.Inventory{
		ProductID: int(request.ProductId),
		VendorID:  int(request.VendorId),
		Quantity:  int(request.Quantity),
	}
}
