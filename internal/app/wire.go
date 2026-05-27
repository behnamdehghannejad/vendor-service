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
	"github.com/behnamdehghannejad/vendorservice/internal/port"
	"github.com/behnamdehghannejad/vendorservice/internal/service"
	"github.com/behnamdehghannejad/vendorservice/internal/validator"
	"github.com/gin-gonic/gin"
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

	if err := migrate(cfg.Database); err != nil {
		return
	}

	_, vendorService, _, err := registerServices(cfg.Database)
	if err != nil {
		return
	}

	server := createServer(cfg.App, vendorService)

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

func registerServices(cfg postgres.PostgresConfig) (port.HistoryService, port.VendorService, port.ProductService, error) {
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, nil, nil, apperror.Wrap(err).UnExpected().DebuggingError().Build()
	}

	historyRepository := postgres.NewHistoryRepository(db)
	vendorRepository := postgres.NewVendorRepository(db)
	productRepository := postgres.NewProductRepository(db)

	historyService := service.NewHistoryService(historyRepository)
	vendorService := service.NewVendorService(vendorRepository)
	productService := service.NewProductService(productRepository)

	return historyService, vendorService, productService, nil
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
	vendorService port.VendorService,
) *http.Server {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())

	orderHandler := httphandler.NewVendorHandler(
		vendorService,
		validator.NewVendorValidator(vendorService),
	)

	registerRoutes(router, orderHandler)

	return &http.Server{
		Addr:              getAddress(cfg.Host, cfg.Port),
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
}

func registerRoutes(
	router *gin.Engine,
	vendorHandler *httphandler.VendorHandler,
) {
	router.POST("/api/v1/vendors", vendorHandler.Create)
	router.GET("/api/v1/vendors/:id", vendorHandler.GetById)
	router.DELETE("/api/v1/vendors/{id}", vendorHandler.Delete)
}

func getAddress(host string, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
