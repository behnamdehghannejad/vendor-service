package repository

import (
	"time"

	"github.com/behnamdehghannejad/vendor/internal/domain"
)

func toHistoryEntity(domain *domain.History) *HistoryEntity {
	return &HistoryEntity{
		OrderID:   domain.OrderID,
		PaymentID: domain.PaymentID,
		Quantity:  domain.Quantity,
		ProductID: domain.ProductID,
		VendorID:  domain.VendorID,
		Active:    domain.Active,
		CreatedAt: time.Now(),
	}
}

func toHistoryDomain(history *HistoryEntity) *domain.History {
	return &domain.History{
		ID:        history.ID,
		OrderID:   history.OrderID,
		PaymentID: history.PaymentID,
		Quantity:  history.Quantity,
		ProductID: history.ProductID,
		VendorID:  history.VendorID,
		Status:    history.Status,
		Active:    history.Active,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
	}
}

func toHistoryDomainValue(history HistoryEntity) domain.History {
	return domain.History{
		ID:        history.ID,
		OrderID:   history.OrderID,
		PaymentID: history.PaymentID,
		Quantity:  history.Quantity,
		ProductID: history.ProductID,
		VendorID:  history.VendorID,
		Status:    history.Status,
		Active:    history.Active,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
	}
}

func toProductDomain(product *ProductEntity) *domain.Product {
	return &domain.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func toProductEntity(product *domain.Product) *ProductEntity {
	return &ProductEntity{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func toVendorEntity(vendor *domain.Vendor) *VendorEntity {
	return &VendorEntity{
		ID:        vendor.ID,
		Name:      vendor.Name,
		Code:      vendor.Code,
		Email:     vendor.Email,
		Phone:     vendor.Phone,
		Address:   vendor.Address,
		Active:    vendor.Active,
		CreatedAt: vendor.CreatedAt,
		UpdatedAt: vendor.UpdatedAt,
	}
}

func toVendorDomain(vendor *VendorEntity) *domain.Vendor {
	return &domain.Vendor{
		ID:        vendor.ID,
		Name:      vendor.Name,
		Code:      vendor.Code,
		Email:     vendor.Email,
		Phone:     vendor.Phone,
		Address:   vendor.Address,
		Active:    vendor.Active,
		CreatedAt: vendor.CreatedAt,
		UpdatedAt: vendor.UpdatedAt,
	}
}
