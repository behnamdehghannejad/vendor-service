package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	discount "github.com/behnamdehghannejad/vendorservice/internal/adapter/discount_client"
	"github.com/behnamdehghannejad/vendorservice/internal/handler/httphandler"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/config"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/log"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/metrics"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/behnamdehghannejad/vendorservice/internal/scheduler"
	"github.com/behnamdehghannejad/vendorservice/internal/service"
	"github.com/behnamdehghannejad/vendorservice/internal/validator"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunHttp() {
	os.Setenv("CONFIG_PATH", "./")

	err := log.Initialize()
	if err != nil {
		return
	}

	cfg, err := config.Load()
	if err != nil {
		return
	}

	metrics.Init()
	if err := migrate(cfg.Database); err != nil {
		return
	}

	transactionService, vendorService, productService, inventoryService, err := registerServices(cfg)
	if err != nil {
		return
	}

	server := createServer(
		cfg.App,
		transactionService,
		vendorService,
		productService,
		inventoryService,
	)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	errCh := make(chan error, 1)

	go startServer(server, cfg.App, errCh)

	select {
	case <-stop:
		shutdownServer(server)
	case err := <-errCh:
		log.Fatalf("server failed %v", err)
	}

	shutdownServer(server)
}

func RunScheduler() {
	os.Setenv("CONFIG_PATH", "./")

	err := log.Initialize()
	if err != nil {
		return
	}

	cfg, err := config.Load()
	if err != nil {
		return
	}

	if err := migrate(cfg.Database); err != nil {
		return
	}

	_, _, productService, _, err := registerServices(cfg)
	if err != nil {
		return
	}

	done := make(chan struct{})

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		scheduler := scheduler.New(productService)
		scheduler.Start(done, &wg)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	done <- struct{}{}
	wg.Wait()
}

func migrate(cfg postgres.PostgresConfig) error {
	migrator, err := postgres.NewMigrator(cfg)
	if err != nil {
		return err
	}
	if err := migrator.UP(); err != nil {
		return err
	}
	return nil
}

func registerServices(cfg config.Config) (
	port.TransactionService,
	port.VendorService,
	port.ProductService,
	port.InventoryService,
	error,
) {
	db, err := postgres.New(cfg.Database)
	if err != nil {
		return nil, nil, nil, nil, apperror.Wrap(err).UnExpected().DebuggingError().Build()
	}

	transactionRepository := postgres.NewTransactionRepository(db)
	vendorRepository := postgres.NewVendorRepository(db)
	productRepository := postgres.NewProductRepository(db)
	inventoryRepository := postgres.NewInventoryRepository(db)

	unitOfWorkFactory := postgres.NewUnitOfWorkFactory(db)

	transactionService := service.NewTransactionService(transactionRepository)
	vendorService := service.NewVendorService(vendorRepository)
	productService := service.NewProductService(
		productRepository,
		discount.New(cfg.DiscountClient.URL),
	)
	inventoryService := service.NewInventoryService(inventoryRepository, unitOfWorkFactory)

	return transactionService, vendorService, productService, inventoryService, nil
}

func startServer(server *http.Server, cfg httphandler.HttpConfig, errCh chan<- error) {
	log.Infof("server started on %s", getAddress(cfg.Host, cfg.Port))

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Warningf("server error: %s", err.Error())
		errCh <- err
	}
}

func shutdownServer(server *http.Server) {
	log.Warning("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Warningf("shutdown failed: %s", err.Error())
		return
	}

	log.Warning("server stopped cleanly")
}

func createServer(
	cfg httphandler.HttpConfig,
	transactionService port.TransactionService,
	vendorService port.VendorService,
	productService port.ProductService,
	inventoryService port.InventoryService,
) *http.Server {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(metrics.PrometheusMiddleware())

	transactionHandler, vendorHandler, productHandler, inventoryHandler := registerHandlers(
		transactionService,
		inventoryService,
		vendorService,
		productService,
	)

	registerRoutes(
		router,
		transactionHandler,
		vendorHandler,
		productHandler,
		inventoryHandler,
	)

	return &http.Server{
		Addr:              getAddress(cfg.Host, cfg.Port),
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
}

func registerHandlers(
	transactionService port.TransactionService,
	inventoryService port.InventoryService,
	vendorService port.VendorService,
	productService port.ProductService,
) (
	*httphandler.Transaction,
	*httphandler.Vendor,
	*httphandler.Product,
	*httphandler.Inventory,
) {
	transactionHandler := httphandler.NewTransactionHandler(
		transactionService,
	)

	vendorHandler := httphandler.NewVendorHandler(
		vendorService,
		validator.NewVendor(vendorService),
	)

	productHandler := httphandler.NewProductHandler(
		productService,
		validator.NewProduct(productService),
	)

	inventoryHandler := httphandler.NewInventory(
		inventoryService,
		validator.NewInventory(inventoryService),
	)

	return transactionHandler, vendorHandler, productHandler, inventoryHandler
}

func registerRoutes(
	router *gin.Engine,
	transactionHandler *httphandler.Transaction,
	vendorHandler *httphandler.Vendor,
	productHandler *httphandler.Product,
	inventoryHandler *httphandler.Inventory,
) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"pong": true})
	})
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/api/v1/inventories", inventoryHandler.Search)
	router.GET(
		"/api/v1/inventories/:vpIDs",
		inventoryHandler.GetInventory,
	)
	router.PUT(
		"/api/v1/inventories/:vpIDs",
		inventoryHandler.Upsert,
	)
	router.POST(
		"/api/v1/inventories/:vpIDs/reserve",
		inventoryHandler.Reserve,
	)

	router.POST("/api/v1/vendors", vendorHandler.Create)
	router.GET("/api/v1/vendors/:id", vendorHandler.GetById)
	router.DELETE("/api/v1/vendors/:id", vendorHandler.Delete)
	router.GET("/api/v1/vendors", vendorHandler.Filter)

	router.POST("/api/v1/products", productHandler.Create)
	router.GET("/api/v1/products/:id", productHandler.GetById)
	router.GET("api/v1/products", productHandler.Filter)
	router.PATCH("api/v1/products/:id", productHandler.Update)

	router.GET("/api/v1/transactions", transactionHandler.Search)
}

func getAddress(host string, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
