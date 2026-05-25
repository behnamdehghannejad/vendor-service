package app

import (
	"log"
	"net"
	"order-service/internal/handler/grpc"

	pb "order-service/proto/generate"

	"google.golang.org/grpc"
)

type Services struct {
	User pb.OrderServiceServer
}

func RunGrpcServer(grpcAddress string, orderHandler *handler.OrderGrpcHandler) {
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderHandler)

	log.Println("gRPC server running on", grpcAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
