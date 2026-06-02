package repository

import (
	"testing"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

func TestVendorRepository_Create(t *testing.T) {
	vendor := domain.Vendor{
		Name:    "Test Vendor",
		Code:    "TV001",
		Email:   "test@example.com",
		Phone:   "123456789",
		Address: "Test Address",
		Active:  true,
	}

	ID, err := vendorRepo.Create(vendor)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	vendorRepo.DeleteVendorsByIDs(ID)
}

func TestVendorRepository_FindByID(t *testing.T) {
	vendor := domain.Vendor{
		Name:   "Find Me",
		Code:   "FIND001",
		Email:  "find@example.com",
		Active: true,
	}

	ID, err := vendorRepo.Create(vendor)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer vendorRepo.DeleteVendorsByIDs(ID)

	found, err := vendorRepo.FindByID(ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if found.Name != "Find Me" {
		t.Fatalf("expected name 'Find Me', got %s", found.Name)
	}
}

func TestVendorRepository_Filter(t *testing.T) {
	firstID, _ := vendorRepo.Create(domain.Vendor{
		Name:   "Apple Vendor",
		Code:   "APL001",
		Active: true,
	})

	secondID, _ := vendorRepo.Create(domain.Vendor{
		Name:   "Banana Vendor",
		Code:   "BAN001",
		Active: false,
	})

	defer vendorRepo.DeleteVendorsByIDs(firstID, secondID)

	active := true

	result, err := vendorRepo.Filter(domain.SearchVendor{
		IsActive: &active,
	})
	if err != nil {
		t.Fatalf("filter failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("expected at least 1 active vendor")
	}
}

func TestVendorRepository_Update(t *testing.T) {
	v := domain.Vendor{
		Name:   "Old Name",
		Code:   "UPD001",
		Active: true,
	}

	ID, err := vendorRepo.Create(v)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	defer vendorRepo.DeleteVendorsByIDs(ID)

	v.Name = "New Name"
	v.ID = ID

	if err := vendorRepo.Update(v); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	updated, err := vendorRepo.FindByID(ID)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}

	if updated.Name != "New Name" {
		t.Fatalf("expected updated name, got %s", updated.Name)
	}
}

func TestVendorRepository_SoftDelete(t *testing.T) {
	v := domain.Vendor{
		Name:   "To Delete",
		Code:   "DEL001",
		Active: true,
	}

	ID, err := vendorRepo.Create(v)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	defer vendorRepo.DeleteVendorsByIDs(ID)

	if err := vendorRepo.SoftDelete(ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	found, err := vendorRepo.FindByID(ID)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}

	if found.Active != false {
		t.Fatalf("expected inactive vendor after delete")
	}
}
