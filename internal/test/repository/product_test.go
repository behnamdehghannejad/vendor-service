package repository

import (
	"testing"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

func TestProductRepository_Create(t *testing.T) {
	categoryID, err := categoryRepo.Create(domain.Category{
		Name:   "Test Category",
		Active: true,
	})
	if err != nil {
		t.Fatalf("create category failed: %v", err)
	}

	p := domain.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Active:      true,
		CategoryID:  categoryID,
	}

	id, err := productRepo.Create(p)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	defer productRepo.DeleteProductsByIDs(id)
	defer categoryRepo.Delete(categoryID)

	found, err := productRepo.FindById(id)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}

	if found.Name != p.Name {
		t.Fatalf("expected %s, got %s", p.Name, found.Name)
	}
}

func TestProductRepository_FindById(t *testing.T) {
	p := domain.Product{
		Name:   "Find Me",
		Active: true,
	}

	id, err := productRepo.Create(p)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer productRepo.DeleteProductsByIDs(id)

	found, err := productRepo.FindById(id)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}

	if found.Name != "Find Me" {
		t.Fatalf("expected 'Find Me', got %s", found.Name)
	}
}

func TestProductRepository_Filter(t *testing.T) {
	firstID, _ := productRepo.Create(domain.Product{
		Name:   "Apple Product",
		Active: true,
	})

	secondID, _ := productRepo.Create(domain.Product{
		Name:   "Banana Product",
		Active: false,
	})

	defer productRepo.DeleteProductsByIDs(firstID, secondID)

	result, err := productRepo.Filter(domain.SearchProduct{
		SearchName: "Apple",
	})
	if err != nil {
		t.Fatalf("filter failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("expected at least 1 product")
	}
}

func TestProductRepository_Update(t *testing.T) {
	p := domain.Product{
		Name:   "Old Name",
		Active: true,
	}

	id, err := productRepo.Create(p)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer productRepo.DeleteProductsByIDs(id)

	p.ID = id
	p.Name = "New Name"

	if err := productRepo.Update(p); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	updated, err := productRepo.FindById(id)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}

	if updated.Name != "New Name" {
		t.Fatalf("expected 'New Name', got %s", updated.Name)
	}
}

func TestProductRepository_SoftDelete(t *testing.T) {
	p := domain.Product{
		Name:   "To Delete",
		Active: true,
	}

	id, err := productRepo.Create(p)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer productRepo.DeleteProductsByIDs(id)

	if err := productRepo.SoftDelete(id); err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	found, err := productRepo.FindById(id)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}

	if found.Active != false {
		t.Fatalf("expected inactive product after delete")
	}
}

func TestProductRepository_FindByCategoryId(t *testing.T) {
	categoryID, err := categoryRepo.Create(domain.Category{
		Name:   "Electronics",
		Active: true,
	})
	if err != nil {
		t.Fatalf("create category failed: %v", err)
	}

	productID, err := productRepo.Create(domain.Product{
		Name:       "Laptop",
		CategoryID: categoryID,
		Active:     true,
	})
	if err != nil {
		t.Fatalf("create product failed: %v", err)
	}

	defer productRepo.DeleteProductsByIDs(productID)
	defer categoryRepo.Delete(categoryID)

	products, err := productRepo.FindByCategoryId(categoryID)
	if err != nil {
		t.Fatalf("find by category failed: %v", err)
	}

	if len(products) == 0 {
		t.Fatal("expected at least one product")
	}

	if products[0].CategoryID != categoryID {
		t.Fatalf("expected category %d, got %d", categoryID, products[0].CategoryID)
	}
}
