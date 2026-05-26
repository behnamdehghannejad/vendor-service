package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "vendor-service/proto/generate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunHTTPGateway(grpcAddress string, httpPort int) {
	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	register := register(grpcAddress, mux, dialOptions...)
	register(pb.RegisterHistoryServiceHandlerFromEndpoint)
	register(pb.RegisterProductServiceHandlerFromEndpoint)
	register(pb.RegisterVendorServiceHandlerFromEndpoint)

	log.Printf("HTTP Gateway running on :%d\n", httpPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux); err != nil {
		log.Fatal(err)
	}
}

func register(grpcAddress string, mux *runtime.ServeMux, dialOptions ...grpc.DialOption) func(registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error) {
	const retries = 1
	const delayMs = 200
	ctx := context.Background()

	return func(registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error) {
		var err error
		for i := 0; i < retries; i++ {
			err = registerFunc(ctx, mux, grpcAddress, dialOptions)
			if err != nil {
				log.Printf("Waiting for gRPC server to start (%d/%d): %v", i+1, retries, err)
			}
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
		if err != nil {
			log.Fatal("Cannot register service:", err)
		}
	}
}
