package app

import (
	"fmt"
	"log"
	"net/http"
	handler2 "vendor-service/internal/handler/http"
	"vendor-service/internal/service"
)

func runHttp(cfg *Config, vendorService service.VendorService, productService service.ProductService, historyService service.HistoryService, inventoryService service.InventoryService, orderService *service.OrderServiceImpl) {
	mux := http.NewServeMux()
	handelVendorRequests(mux, handler2.NewVendorHandler(vendorService))
	handelProductRequests(mux, handler2.NewProductHandler(productService))
	handelHistoryRequests(mux, handler2.NewHistoryHandler(historyService))
	handelInventoryRequests(mux, handler2.NewInventoryHandler(inventoryService))
	handelOrderRequests(mux, handler2.NewOrderHandler(orderService))
	port := cfg.App.HttpPort
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	log.Println("HTTP server running on", port)
}

func handelOrderRequests(mux *http.ServeMux, handler *handler2.OrderHandler) {
	mux.HandleFunc("POST /api/order/create", handler.ManageOrders)
	mux.HandleFunc("POST /api/vendor/payment", handler.AcceptOrdersPayment)
}

func handelVendorRequests(mux *http.ServeMux, handler *handler2.VendorHandler) {
	mux.HandleFunc("POST /api/vendor/create", handler.Create)
	mux.HandleFunc("POST /api/vendor/update", handler.Update)
	mux.HandleFunc("GET /api/vendor/{id}", handler.GetById)
	mux.HandleFunc("GET /api/vendor/delete/{id}", handler.Delete)
	mux.HandleFunc("GET /api/vendor/code/{id}", handler.GetByCode)
}

func handelProductRequests(mux *http.ServeMux, handler *handler2.ProductHandler) {
	mux.HandleFunc("POST /api/product/create", handler.Create)
	mux.HandleFunc("PUT /api/product/update", handler.Update)
	mux.HandleFunc("GET /api/product/{id}", handler.GetById)
	mux.HandleFunc("DELETE /api/product/delete/{id}", handler.Delete)
}

func handelHistoryRequests(mux *http.ServeMux, handler *handler2.HistoryHandler) {
	mux.HandleFunc("POST /api/history/create", handler.Create)
	mux.HandleFunc("GET /api/history/order/{id}", handler.GetByOrderID)
	mux.HandleFunc("GET /api/history/vendor/{id}", handler.GetByVendorID)
	mux.HandleFunc("GET /api/history/payment/{id}", handler.GetByPaymentID)
	mux.HandleFunc("GET /api/history/product/{id}", handler.GetByProductID)
	mux.HandleFunc("GET /api/history/actives", handler.GetByIsActive)
	mux.HandleFunc("POST /api/history/status", handler.GetByStatus)
	mux.HandleFunc("DELETE /api/history/delete/{id}", handler.Delete)

}

func handelInventoryRequests(mux *http.ServeMux, handler *handler2.InventoryHandler) {
	mux.HandleFunc("POST /api/inventory/api/orders", handler.AddProductsToVendor)
}
