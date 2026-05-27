package handler

import (
	"context"
	"time"

	"github.com/behnamdehghannejad/vendor/internal/domain"
	"github.com/behnamdehghannejad/vendor/internal/service"
	pb "github.com/behnamdehghannejad/vendor/proto/generate"

	"google.golang.org/protobuf/types/known/emptypb"
)

type VendorGrpcHandler struct {
	pb.UnimplementedVendorServiceServer
	service service.VendorService
}

func NewVendorGrpcHandler(service service.VendorService) *VendorGrpcHandler {
	return &VendorGrpcHandler{
		service: service,
	}
}

func (handler *VendorGrpcHandler) Add(ctx context.Context, request *pb.CreateVendorRequest) (*emptypb.Empty, error) {
	requestDto := vendorProtoToDomain(request)

	if err := handler.service.Create(requestDto); err != nil {
		return nil, err
	}

	return nil, nil
}

func (handler *VendorGrpcHandler) Update(ctx context.Context, request *pb.UpdateVendorRequest) (*emptypb.Empty, error) {
	vendor := toVendorDomain(request.Vendor)

	if err := handler.service.Update(vendor); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (handler *VendorGrpcHandler) Delete(ctx context.Context, request *pb.DeleteVendorRequest) (*emptypb.Empty, error) {
	if err := handler.service.Delete(int(request.Id)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (handler *VendorGrpcHandler) FindByID(ctx context.Context, request *pb.GetVendorRequest) (*pb.VendorResponse, error) {
	vendor, err := handler.service.FindByID(int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.VendorResponse{
		Vendor: toVendorProto(vendor),
	}, nil
}

func (handler *VendorGrpcHandler) FindByCode(ctx context.Context, request *pb.GetCodeRequest) (*pb.VendorResponse, error) {
	vendor, err := handler.service.FindByCode(request.Code)
	if err != nil {
		return nil, err
	}

	return &pb.VendorResponse{
		Vendor: toVendorProto(vendor),
	}, nil
}

func toVendorDomain(vendor *pb.Vendor) *domain.Vendor {
	return &domain.Vendor{
		ID:      int(vendor.Id),
		Code:    vendor.Code,
		Name:    vendor.Name,
		Email:   vendor.Email,
		Phone:   vendor.Phone,
		Address: vendor.Address,
		Active:  vendor.Active,
	}
}

func vendorProtoToDomain(request *pb.CreateVendorRequest) *domain.Vendor {
	return &domain.Vendor{
		Code:    request.Code,
		Name:    request.Name,
		Email:   request.Email,
		Phone:   request.Phone,
		Address: request.Address,
		Active:  true,
	}
}

func toVendorProto(vendor *domain.Vendor) *pb.Vendor {
	return &pb.Vendor{
		Id:        int32(vendor.ID),
		Code:      vendor.Code,
		Name:      vendor.Name,
		Email:     vendor.Email,
		Phone:     vendor.Phone,
		Address:   vendor.Address,
		Active:    vendor.Active,
		CreatedAt: vendor.CreatedAt.Format(time.RFC3339),
		UpdatedAt: vendor.UpdatedAt.Format(time.RFC3339),
	}
}
