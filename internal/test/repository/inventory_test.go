package repository

import (
	"fmt"
	"testing"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

func TestInventoryRepository_Upsert_Update(t *testing.T) {
	productID, _ := productRepo.Create(domain.Product{})
	defer productRepo.DeleteProductsByIDs(productID)
	vendorID, _ := vendorRepo.Create(domain.Vendor{})
	defer vendorRepo.DeleteVendorsByIDs(vendorID)

	fmt.Println(productID, vendorID)
	inv := domain.Inventory{
		VendorID:  vendorID,
		ProductID: productID,
		Quantity:  50,
		Reserved:  0,
		V:         0,
	}

	if err := inventoryRepo.Upsert(inv); err != nil {
		t.Fatalf("create failed: %v", err)
	}
	defer inventoryRepo.DeleteInventoriesByID(vendorID, productID)

	inv.Quantity = 80
	inv.V = 0

	if err := inventoryRepo.Upsert(inv); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	inventory, err := inventoryRepo.GetInventory(vendorID, productID)
	if err != nil {
		t.Fatalf("error to get the inventory the problem is %v", err)
	}

	if inventory.Quantity != inv.Quantity {
		t.Fatalf(
			"quantity mismatch: expected %d, got %d (vendor_id=%d, product_id=%d)",
			inv.Quantity,
			inventory.Quantity,
			inv.VendorID,
			inv.ProductID,
		)
	}
}

func TestInventoryRepository_IncreaseReserveInventory(t *testing.T) {
	productID, _ := productRepo.Create(domain.Product{})
	defer productRepo.DeleteProductsByIDs(productID)
	vendorID, _ := vendorRepo.Create(domain.Vendor{})
	defer vendorRepo.DeleteVendorsByIDs(vendorID)

	inv := domain.Inventory{
		VendorID:  vendorID,
		ProductID: productID,
		Quantity:  100,
		Reserved:  10,
		V:         0,
	}
	defer inventoryRepo.DeleteInventoriesByID(vendorID, productID)

	if err := inventoryRepo.Upsert(inv); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	req := domain.RequestReserve{
		VendorID:  3,
		ProductID: 3,
		Reserved:  5,
	}

	if err := inventoryRepo.IncreaseReserveInventory(req); err != nil {
		t.Fatalf("increase reserve failed: %v", err)
	}
}

func TestInventoryRepository_GetInventory(t *testing.T) {
	productID, _ := productRepo.Create(domain.Product{})
	defer productRepo.DeleteProductsByIDs(productID)
	vendorID, _ := vendorRepo.Create(domain.Vendor{})
	defer vendorRepo.DeleteVendorsByIDs(vendorID)

	inv := domain.Inventory{
		VendorID:  vendorID,
		ProductID: productID,
		Quantity:  200,
		Reserved:  0,
		V:         0,
	}

	if err := inventoryRepo.Upsert(inv); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	defer inventoryRepo.DeleteInventoriesByID(vendorID, productID)

	found, err := inventoryRepo.GetInventory(vendorID, productID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if found.Quantity != 200 {
		t.Fatalf("expected 200, got %v", found.Quantity)
	}
}

func TestInventoryRepository_Filter(t *testing.T) {
	firstProductID, _ := productRepo.Create(domain.Product{})
	defer productRepo.DeleteProductsByIDs(firstProductID)
	firstVendorID, _ := vendorRepo.Create(domain.Vendor{})
	defer vendorRepo.DeleteVendorsByIDs(firstVendorID)

	_ = inventoryRepo.Upsert(domain.Inventory{
		VendorID:  firstVendorID,
		ProductID: firstProductID,
		Quantity:  10,
		V:         0,
	})
	defer inventoryRepo.DeleteInventoriesByID(firstVendorID, firstProductID)

	secondProductID, _ := productRepo.Create(domain.Product{})
	defer productRepo.DeleteProductsByIDs(secondProductID)
	secondVendorID, _ := vendorRepo.Create(domain.Vendor{})
	defer vendorRepo.DeleteVendorsByIDs(secondVendorID)
	_ = inventoryRepo.Upsert(domain.Inventory{
		VendorID:  secondVendorID,
		ProductID: secondProductID,
		Quantity:  20,
		V:         0,
	})

	defer inventoryRepo.DeleteInventoriesByID(secondVendorID, secondProductID)

	vendorID := firstVendorID

	result, err := inventoryRepo.Filter(domain.SearchInventory{
		VendorID: &vendorID,
	})
	if err != nil {
		t.Fatalf("filter failed: %v", err)
	}

	if len(result) != 1 {
		t.Fatal("expected at least 1 inventory")
	}
}

func TestInventoryRepository_AcceptReserve(t *testing.T) {
	productID, _ := productRepo.Create(domain.Product{})
	defer productRepo.DeleteProductsByIDs(productID)
	vendorID, _ := vendorRepo.Create(domain.Vendor{})
	defer vendorRepo.DeleteVendorsByIDs(vendorID)

	inv := domain.Inventory{
		VendorID:  vendorID,
		ProductID: productID,
		Quantity:  100,
		Reserved:  20,
		V:         0,
	}

	if err := inventoryRepo.Upsert(inv); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	defer inventoryRepo.DeleteInventoriesByID(inv.VendorID, inv.ProductID)

	final := domain.FinalizeReservation{
		VendorID:  vendorID,
		ProductID: productID,
		Reserve:   10,
	}

	if err := inventoryRepo.AcceptReserve(final); err != nil {
		t.Fatalf("accept reserve failed: %v", err)
	}

	updated, err := inventoryRepo.GetInventory(vendorID, productID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if updated.Quantity != 90 {
		t.Fatalf("expected 90, got %v", updated.Quantity)
	}

	if updated.Reserved != 10 {
		t.Fatalf("expected 10, got %v", updated.Reserved)
	}
}

func TestInventoryRepository_RejectReserve(t *testing.T) {
	productID, _ := productRepo.Create(domain.Product{})
	defer productRepo.DeleteProductsByIDs(productID)
	vendorID, _ := vendorRepo.Create(domain.Vendor{})
	defer vendorRepo.DeleteVendorsByIDs(vendorID)

	inv := domain.Inventory{
		VendorID:  vendorID,
		ProductID: productID,
		Quantity:  100,
		Reserved:  30,
		V:         0,
	}

	if err := inventoryRepo.Upsert(inv); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	defer inventoryRepo.DeleteInventoriesByID(inv.VendorID, inv.ProductID)

	final := domain.FinalizeReservation{
		VendorID:  vendorID,
		ProductID: productID,
		Reserve:   10,
	}

	if err := inventoryRepo.RejectReserve(final); err != nil {
		t.Fatalf("reject reserve failed: %v", err)
	}

	updated, err := inventoryRepo.GetInventory(vendorID, productID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if updated.Reserved != 20 {
		t.Fatalf("expected 20, got %v", updated.Reserved)
	}
}

func TestInventoryRepository_UpdateInventoryDiscountPercentages(t *testing.T) {
	productID, _ := productRepo.Create(domain.Product{})
	defer productRepo.DeleteProductsByIDs(productID)

	firstVendorID, _ := vendorRepo.Create(domain.Vendor{
		Code: "1",
	})
	defer vendorRepo.DeleteVendorsByIDs(firstVendorID)

	secondVendorID, _ := vendorRepo.Create(domain.Vendor{
		Code: "2",
	})
	defer vendorRepo.DeleteVendorsByIDs(secondVendorID)

	thirdVendorID, _ := vendorRepo.Create(domain.Vendor{
		Code: "3",
	})
	defer vendorRepo.DeleteVendorsByIDs(thirdVendorID)

	inventoryDiscountPercentages := make([]domain.InventoryDiscountPercentage, 0, 3)

	discountPercentage := 3

	for _, inventory := range []domain.Inventory{
		{ProductID: productID, VendorID: firstVendorID, Quantity: 20},
		{ProductID: productID, VendorID: secondVendorID, Quantity: 30},
		{ProductID: productID, VendorID: thirdVendorID, Quantity: 40},
	} {
		err := inventoryRepo.Upsert(inventory)
		if err != nil {
			t.Fatalf("error creating inventory: %v", err)
		}
		defer inventoryRepo.DeleteInventoriesByID(inventory.VendorID, inventory.ProductID)

		inventoryDiscountPercentages = append(
			inventoryDiscountPercentages,
			domain.InventoryDiscountPercentage{
				ProductID:          inventory.ProductID,
				VendorID:           inventory.VendorID,
				DiscountPercentage: float64(discountPercentage),
			},
		)

		discountPercentage++
	}

	err := inventoryRepo.UpdateProductDiscountPercentages(inventoryDiscountPercentages)
	if err != nil {
		t.Fatalf("error updating discounts: %v", err)
	}

	for _, expected := range inventoryDiscountPercentages {
		inventory, err := inventoryRepo.GetInventory(expected.VendorID, expected.ProductID)
		if err != nil {
			t.Fatalf("error finding product: %v", err)
		}

		if inventory.DiscountPercentage != expected.DiscountPercentage {
			t.Fatalf(
				"expected %f but got %f",
				expected.DiscountPercentage,
				inventory.DiscountPercentage,
			)
		}
	}
}
