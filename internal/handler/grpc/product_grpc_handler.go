package handler

import (
	"context"
	"time"

	"github.com/behnamdehghannejad/vendor/internal/domain"
	"github.com/behnamdehghannejad/vendor/internal/service"
	pb "github.com/behnamdehghannejad/vendor/proto/generate"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductGrpcHandler struct {
	pb.UnimplementedProductServiceServer
	service service.ProductService
}

func NewProductGrpcHandler(service service.ProductService) *ProductGrpcHandler {
	return &ProductGrpcHandler{
		service: service,
	}
}

func (handler *ProductGrpcHandler) Add(ctx context.Context, request *pb.CreateProductRequest) (*emptypb.Empty, error) {
	requestDto := toProductDomain(request)

	if err := handler.service.Create(requestDto); err != nil {
		return nil, err
	}

	return nil, nil
}

func (handler *ProductGrpcHandler) Update(ctx context.Context, request *pb.UpdateProductRequest) (*emptypb.Empty, error) {
	product := toProductProtoDomain(request.Product)

	if err := handler.service.Update(product); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (handler *ProductGrpcHandler) Delete(ctx context.Context, request *pb.DeleteProductRequest) (*emptypb.Empty, error) {
	if err := handler.service.Delete(int(request.Id)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (handler *ProductGrpcHandler) FindById(ctx context.Context, request *pb.GetProductRequest) (*pb.ProductResponse, error) {
	product, err := handler.service.FindById(int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Product: toProductProto(product),
	}, nil
}

func toProductDomain(request *pb.CreateProductRequest) *domain.Product {
	return &domain.Product{
		Name:        request.Name,
		Description: request.Description,
		Active:      true,
	}
}

func toProductProtoDomain(product *pb.Product) *domain.Product {
	return &domain.Product{
		ID:          int(product.Id),
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
	}
}

func toProductProto(product *domain.Product) *pb.Product {
	return &pb.Product{
		Id:          int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}
}
