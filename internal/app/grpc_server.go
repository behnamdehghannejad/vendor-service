package app

import (
	"log"
	"net"

	handler "github.com/behnamdehghannejad/vendor/internal/handler/grpc"

	pb "github.com/behnamdehghannejad/vendor/proto/generate"

	"google.golang.org/grpc"
)

type Services struct {
	Vendor  pb.VendorServiceServer
	Product pb.ProductServiceServer
	History pb.HistoryServiceServer
}

func RunGrpcServer(grpcAddress string,
	vendorGrpcHandler handler.VendorGrpcHandler,
	productGrpcHandler handler.ProductGrpcHandler,
	historyGrpcHandler handler.HistoryGrpcHandler,
) {
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVendorServiceServer(grpcServer, &vendorGrpcHandler)
	pb.RegisterProductServiceServer(grpcServer, &productGrpcHandler)
	pb.RegisterHistoryServiceServer(grpcServer, &historyGrpcHandler)

	log.Println("gRPC server running on", grpcAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
