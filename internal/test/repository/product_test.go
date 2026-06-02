package repository

import (
	"testing"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

func TestProductRepository_Create(t *testing.T) {
	p := domain.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Active:      true,
	}

	ID, err := productRepo.Create(p)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	productRepo.DeleteProductsByIDs(ID)
}

func TestProductRepository_FindById(t *testing.T) {
	p := domain.Product{
		Name:   "Find Me",
		Active: true,
	}

	ID, err := productRepo.Create(p)
	if err != nil {
		t.Fatalf("filter failed: %v", err)
	}
	defer productRepo.DeleteProductsByIDs(ID)

	found, err := productRepo.FindById(ID)
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
	ID, err := productRepo.Create(p)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer productRepo.DeleteProductsByIDs(ID)

	p.ID = ID
	p.Name = "New Name"

	if err := productRepo.Update(p); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	updated, err := productRepo.FindById(ID)
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

	ID, err := productRepo.Create(p)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer productRepo.DeleteProductsByIDs(ID)

	if err := productRepo.SoftDelete(ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	found, err := productRepo.FindById(ID)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}

	if found.Active != false {
		t.Fatalf("expected inactive product after delete")
	}
}

func TestProductRepository_UpdateProductDiscountPercentages(t *testing.T) {
	productDiscountPercentages := make([]domain.ProductDiscountPercentage, 0, 3)
	ids := make([]int, 0, 3)

	discountPercentage := 3

	for _, product := range []domain.Product{
		{Name: "book1", Description: "1"},
		{Name: "book2", Description: "2"},
		{Name: "book3", Description: "3"},
	} {
		ID, err := productRepo.Create(product)
		if err != nil {
			t.Fatalf("error creating product: %v", err)
		}
		defer productRepo.DeleteProductsByIDs(ID)

		ids = append(ids, ID)

		productDiscountPercentages = append(
			productDiscountPercentages,
			domain.ProductDiscountPercentage{
				ProductID:          ID,
				DiscountPercentage: float64(discountPercentage),
			},
		)

		discountPercentage++
	}

	err := productRepo.UpdateProductDiscountPercentages(productDiscountPercentages)
	if err != nil {
		t.Fatalf("error updating discounts: %v", err)
	}

	for _, expected := range productDiscountPercentages {
		product, err := productRepo.FindById(expected.ProductID)
		if err != nil {
			t.Fatalf("error finding product: %v", err)
		}

		if product.DiscountPercentage != expected.DiscountPercentage {
			t.Fatalf(
				"expected %f but got %f",
				expected.DiscountPercentage,
				product.DiscountPercentage,
			)
		}
	}
}
