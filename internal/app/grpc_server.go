package app

import (
	"log"
	"net"
	"vendor-service/internal/handler/grpc"

	pb "vendor-service/proto/generate"

	"google.golang.org/grpc"
)

type Services struct {
	Vendor  pb.VendorServiceServer
	Product pb.ProductServiceServer
	History pb.HistoryServiceServer
}

func RunGrpcServer(
	grpcAddress string,
	vendorGrpcHandler handler.VendorGrpcHandler,
	productGrpcHandler handler.ProductGrpcHandler,
	historyGrpcHandler handler.HistoryGrpcHandler,
	inventoryGrpcHandler handler.InventoryGrpcHandler,
	orderGrpcHandler handler.OrderGrpcHandler,
) {
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVendorServiceServer(grpcServer, &vendorGrpcHandler)
	pb.RegisterProductServiceServer(grpcServer, &productGrpcHandler)
	pb.RegisterHistoryServiceServer(grpcServer, &historyGrpcHandler)
	pb.RegisterInventoryServiceServer(grpcServer, &inventoryGrpcHandler)
	pb.RegisterOrderServiceServer(grpcServer, &orderGrpcHandler)

	log.Println("gRPC server running on", grpcAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
