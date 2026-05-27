package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	handler "github.com/behnamdehghannejad/vendor/internal/handler/grpc"
	handler2 "github.com/behnamdehghannejad/vendor/internal/handler/http"
	"github.com/behnamdehghannejad/vendor/internal/infra/repository"
	"github.com/behnamdehghannejad/vendor/internal/service"
)

func Start() error {
	cfg := LoadConfig()
	postgres, err := NewPostgres(cfg)
	if err != nil {
		return err
	}

	err = postgres.AutoMigrate(&repository.ProductEntity{})
	err = postgres.AutoMigrate(&repository.VendorEntity{})
	// err = postgres.AutoMigrate(&repository.InventoryEntity{})
	err = postgres.AutoMigrate(&repository.HistoryEntity{})

	vendorRepository := repository.NewVendorRepositoryImpl(postgres)
	productRepository := repository.NewProductRepositoryImpl(postgres)
	historyRepository := repository.NewHistoryRepositoryImpl(postgres)
	vendorService := service.NewVendorService(vendorRepository)
	productService := service.NewProductService(productRepository)
	historyService := service.NewHistoryService(historyRepository)

	go runHttp(cfg, vendorService, productService, historyService)
	time.Sleep(100 * time.Millisecond)

	runServer(cfg, vendorService, productService, historyService)

	return err
}

func runHttp(cfg *Config, vendorService service.VendorService, productService service.ProductService, historyService service.HistoryService) {
	mux := http.NewServeMux()
	handelVendorRequests(mux, handler2.NewVendorHandler(vendorService))
	handelProductRequests(mux, handler2.NewProductHandler(productService))
	handelHistoryRequests(mux, handler2.NewHistoryHandler(historyService))
	port := cfg.App.HttpPort
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	log.Println("HTTP server running on", port)
}

func handelVendorRequests(mux *http.ServeMux, orderHandler *handler2.VendorHandler) {
	mux.HandleFunc("POST /vendor/api/orders", orderHandler.Create)
	mux.HandleFunc("GET /vendor/api/orders/{id}", orderHandler.GetById)
	mux.HandleFunc("GET /vendor/api/orders/delete/{id}", orderHandler.Delete)
	// mux.HandleFunc("GET /vendor/api/orders/delete/{id}", orderHandler.GetById)
	// mux.HandleFunc("GET /vendor/api/orders/delete/{id}", orderHandler.GetByCode)
}

func handelProductRequests(mux *http.ServeMux, orderHandler *handler2.ProductHandler) {
	// mux.HandleFunc("POST /vendor/api/orders", orderHandler.Create)
	// mux.HandleFunc("GET /vendor/api/orders/{id}", orderHandler.GetById)
	// mux.HandleFunc("GET /vendor/api/orders/delete/{id}", orderHandler.Delete)
	// mux.HandleFunc("GET /vendor/api/orders/delete/{id}", orderHandler.Update)
}

func handelHistoryRequests(mux *http.ServeMux, orderHandler *handler2.HistoryHandler) {
	// mux.HandleFunc("POST /vendor/api/orders", orderHandler.Create)
	// mux.HandleFunc("POST /vendor/api/orders", orderHandler.GetByOrderID)
	// mux.HandleFunc("POST /vendor/api/orders", orderHandler.GetByVendorID)
	// mux.HandleFunc("POST /vendor/api/orders", orderHandler.GetByPaymentID)
	// mux.HandleFunc("POST /vendor/api/orders", orderHandler.GetByProductID)
	// mux.HandleFunc("POST /vendor/api/orders", orderHandler.GetByIsActive)
	// mux.HandleFunc("POST /vendor/api/orders", orderHandler.GetByStatus)
	// mux.HandleFunc("GET /vendor/api/orders/delete/{id}", orderHandler.Delete)
}

func runServer(cfg *Config, vendorService service.VendorService, productService service.ProductService, historyService service.HistoryService) {
	vendorGrpcHandler := handler.NewVendorGrpcHandler(vendorService)
	productGrpcHandler := handler.NewProductGrpcHandler(productService)
	historyGrpcHandler := handler.NewHistoryGrpcHandler(historyService)
	grpcAddress := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)

	// Start HTTP Gateway in goroutine
	go RunHTTPGateway(grpcAddress, cfg.App.GatewayPort)
	time.Sleep(100 * time.Millisecond)
	// Start gRPC server in main thread (blocking)
	RunGrpcServer(grpcAddress, *vendorGrpcHandler, *productGrpcHandler, *historyGrpcHandler)
}
