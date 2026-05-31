package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/handler/httphandler"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/config"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/log"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/metrics"
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/behnamdehghannejad/vendorservice/internal/service"
	"github.com/behnamdehghannejad/vendorservice/internal/validator"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run() {
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

	historyService, vendorService, productService, inventoryService, err := registerServices(cfg.Database)
	if err != nil {
		return
	}

	server := createServer(
		cfg.App,
		historyService,
		vendorService,
		productService,
		inventoryService,
	)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go startServer(server, cfg.App)

	<-stop

	shutdownServer(server)
}

func migrate(cfg postgres.PostgresConfig) error {
	migrator := postgres.NewMigrator(cfg)
	if err := migrator.UP(); err != nil {
		return err
	}
	return nil
}

func registerServices(cfg postgres.PostgresConfig) (
	port.HistoryService,
	port.VendorService,
	port.ProductService,
	port.InventoryService,
	error,
) {
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, nil, nil, nil, apperror.Wrap(err).UnExpected().DebuggingError().Build()
	}

	historyRepository := postgres.NewHistoryRepository(db)
	vendorRepository := postgres.NewVendorRepository(db)
	productRepository := postgres.NewProductRepository(db)
	inventoryRepository := postgres.NewInventoryRepository(db)

	historyService := service.NewHistoryService(historyRepository)
	vendorService := service.NewVendorService(vendorRepository)
	productService := service.NewProductService(productRepository)
	inventoryService := service.NewInventoryService(inventoryRepository)

	return historyService, vendorService, productService, inventoryService, nil
}

func startServer(server *http.Server, cfg httphandler.HttpConfig) {
	log.Infof("server started on %s", getAddress(cfg.Host, cfg.Port))

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Warningf("server error: %s", err.Error())
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
	historyService port.HistoryService,
	vendorService port.VendorService,
	productService port.ProductService,
	inventoryService port.InventoryService,
) *http.Server {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(metrics.PrometheusMiddleware())

	historyHandler, vendorHandler, productHandler := registerHandlers(
		historyService,
		inventoryService,
		vendorService,
		productService,
	)

	registerRoutes(
		router,
		historyHandler,
		vendorHandler,
		productHandler,
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
	historyService port.HistoryService,
	inventoryService port.InventoryService,
	vendorService port.VendorService,
	productService port.ProductService,
) (
	*httphandler.History,
	*httphandler.Vendor,
	*httphandler.Product,
) {
	historyHandler := httphandler.NewHistoryHandler(
		historyService,
	)

	vendorHandler := httphandler.NewVendorHandler(
		vendorService,
		validator.NewVendor(vendorService),
	)

	productHandler := httphandler.NewProductHandler(
		productService,
		validator.NewProduct(productService),
	)

	return historyHandler, vendorHandler, productHandler
}

func registerRoutes(
	router *gin.Engine,
	historyHandler *httphandler.History,
	vendorHandler *httphandler.Vendor,
	productHandler *httphandler.Product,
) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"pong": true})
	})
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.POST("/api/v1/vendors", vendorHandler.Create)
	router.GET("/api/v1/vendors/:id", vendorHandler.GetById)
	router.DELETE("/api/v1/vendors/:id", vendorHandler.Delete)
	router.GET("/api/v1/vendors", vendorHandler.Filter)

	router.POST("/api/v1/products", productHandler.Create)
	router.GET("/api/v1/products/:id", productHandler.GetById)
	router.GET("api/v1/products", productHandler.Filter)
	router.PATCH("api/v1/products/:id", productHandler.Update)

	router.GET("/api/v1/histories", historyHandler.Search)
}

func getAddress(host string, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
