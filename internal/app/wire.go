package app

import (
	"fmt"
	"time"
	handler "vendor-service/internal/handler/grpc"
	"vendor-service/internal/infra/repository"
	"vendor-service/internal/service"
)

func Start() error {
	cfg := LoadConfig()
	postgres, err := NewPostgres(cfg)
	if err != nil {
		return err
	}

	err = postgres.AutoMigrate(&repository.ProductEntity{})
	err = postgres.AutoMigrate(&repository.VendorEntity{})
	err = postgres.AutoMigrate(&repository.InventoryEntity{})
	err = postgres.AutoMigrate(&repository.HistoryEntity{})

	vendorRepository := repository.NewVendorRepositoryImpl(postgres)
	productRepository := repository.NewProductRepositoryImpl(postgres)
	historyRepository := repository.NewHistoryRepositoryImpl(postgres)
	inventoryRepository := repository.NewInventoryRepositoryImpl(postgres)

	vendorService := service.NewVendorService(vendorRepository)
	productService := service.NewProductService(productRepository)
	historyService := service.NewHistoryService(historyRepository)
	inventoryService := service.NewInventoryService(inventoryRepository)
	orderService := service.NewOrderService(inventoryService, productService, vendorService, historyService)

	// start HTTP in goroutine
	go runHttp(cfg, vendorService, productService, historyService, inventoryService, orderService)
	time.Sleep(100 * time.Millisecond)
	runServer(cfg, vendorService, productService, historyService, inventoryService, orderService)

	return err
}

func runServer(cfg *Config, vendorService service.VendorService, productService service.ProductService, historyService service.HistoryService, inventoryService *service.InventoryServiceImpl, orderService *service.OrderServiceImpl) {
	vendorGrpcHandler := handler.NewVendorGrpcHandler(vendorService)
	productGrpcHandler := handler.NewProductGrpcHandler(productService)
	historyGrpcHandler := handler.NewHistoryGrpcHandler(historyService)
	inventoryGrpcHandler := handler.NewInventoryGrpcHandler(inventoryService)
	orderGrpcHandler := handler.NewOrderGrpcHandler(orderService)

	grpcAddress := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)

	// start HTTP gateway in goroutine
	go RunHTTPGateway(grpcAddress, cfg.App.GatewayPort)
	time.Sleep(100 * time.Millisecond)
	// start gRPC server in main thread (blocking)
	RunGrpcServer(grpcAddress, *vendorGrpcHandler, *productGrpcHandler, *historyGrpcHandler, *inventoryGrpcHandler, *orderGrpcHandler)
}
