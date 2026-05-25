package app

import (
	"fmt"
	"log"
	"net/http"
	handler "order-service/internal/handler/grpc"
	handler2 "order-service/internal/handler/http"
	"order-service/internal/infra/repository"
	"order-service/internal/service"
	"time"
)

func Start() error {
	cfg := LoadConfig()
	postgres, err := NewPostgres(cfg)
	if err != nil {
		return err
	}

	err = postgres.AutoMigrate(&repository.OrderEntity{})

	orderRepository := repository.NewOrderRepository(postgres)
	orderService := service.NewOrderService(orderRepository)
	orderGrpcHandler := handler.NewOrderGrpcHandler(orderService)
	go runHttp(orderService, cfg)
	time.Sleep(100 * time.Millisecond)
	runServer(cfg, orderGrpcHandler)

	return err
}

func runHttp(orderService *service.OrderServiceImpl, cfg *Config) {
	mux := http.NewServeMux()
	handelRequests(mux, handler2.NewOrderHandler(orderService))
	port := cfg.App.HttpPort
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	log.Println("HTTP server running on", port)
}

func handelRequests(mux *http.ServeMux, orderHandler *handler2.OrderHandler) {
	mux.HandleFunc("POST /oms/api/orders", orderHandler.Create)
	mux.HandleFunc("GET /oms/api/orders/{id}", orderHandler.GetById)
	mux.HandleFunc("GET /oms/api/orders/user/{id}", orderHandler.GetByUserId)
	mux.HandleFunc("POST /oms/api/orders/status/{id}", orderHandler.UpdateStatus)
	mux.HandleFunc("GET /oms/api/orders/all", orderHandler.ListAll)
	mux.HandleFunc("GET /oms/api/orders/delete/{id}", orderHandler.Delete)
}

func runServer(cfg *Config, taskHandler *handler.OrderGrpcHandler) {
	grpcAddress := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)

	// Start HTTP Gateway in goroutine
	go RunHTTPGateway(grpcAddress, cfg.App.GatewayPort)
	time.Sleep(100 * time.Millisecond)
	// Start gRPC server in main thread (blocking)
	RunGrpcServer(grpcAddress, taskHandler)
}
