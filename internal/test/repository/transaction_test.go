package repository

import (
	"testing"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

func TestTransactionRepository_Create(t *testing.T) {
	transaction := domain.Transaction{
		Reserved:  100,
		ProductID: 1,
		VendorID:  1,
		Status:    domain.TRANSACTION_DRAFT,
	}

	err := transactionRepo.Create(transaction)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	transactionRepo.DeleteTransactionsByIDs(transaction.ID)
}

func TestTransactionRepository_GetByID(t *testing.T) {
	tr := domain.Transaction{
		ID:        "1",
		Reserved:  200,
		ProductID: 1,
		VendorID:  1,
		Status:    domain.TRANSACTION_DRAFT,
	}

	if err := transactionRepo.Create(tr); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer transactionRepo.DeleteTransactionsByIDs(tr.ID)

	found, err := transactionRepo.GetByID(tr.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if found.Reserved != 200 {
		t.Fatalf("expected 200, got %v", found.Reserved)
	}
}

func TestTransactionRepository_Filter(t *testing.T) {
	firstTr := domain.Transaction{
		ID:        "1",
		Reserved:  300,
		ProductID: 1,
		VendorID:  1,
		Status:    domain.TRANSACTION_DRAFT,
	}

	secondTr := domain.Transaction{
		ID:        "2",
		Reserved:  400,
		ProductID: 2,
		VendorID:  2,
		Status:    domain.TRANSACTION_DRAFT,
	}
	_ = transactionRepo.Create(firstTr)

	_ = transactionRepo.Create(secondTr)

	defer transactionRepo.DeleteTransactionsByIDs(firstTr.ID, secondTr.ID)

	result, err := transactionRepo.Filter(domain.SearchTransaction{})
	if err != nil {
		t.Fatalf("filter failed: %v", err)
	}

	if len(result) != 2 {
		t.Fatal("expected at least 2 transaction")
	}

	productID := 1
	vendorID := 1
	result, err = transactionRepo.Filter(domain.SearchTransaction{
		ProductID: &productID,
		VendorID:  &vendorID,
	})
	if err != nil {
		t.Fatalf("filter failed: %v", err)
	}

	if len(result) != 1 {
		t.Fatal("expected at least 1 transaction")
	}

	productID = 4
	vendorID = 4
	result, err = transactionRepo.Filter(domain.SearchTransaction{
		ProductID: &productID,
		VendorID:  &vendorID,
	})
	if err != nil {
		t.Fatalf("filter failed: %v", err)
	}

	if len(result) != 0 {
		t.Fatal("expected at least 0 transaction")
	}
}

func TestTransactionRepository_Update(t *testing.T) {
	tr := domain.Transaction{
		ID:        "1",
		Reserved:  500,
		ProductID: 1,
		VendorID:  1,
		Status:    domain.TRANSACTION_DRAFT,
	}

	if err := transactionRepo.Create(tr); err != nil {
		t.Fatalf("create failed: %v", err)
	}
	defer transactionRepo.DeleteTransactionsByIDs(tr.ID)

	tr.Reserved = 999

	if err := transactionRepo.Update(tr); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	updated, err := transactionRepo.GetByID(tr.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if updated.Reserved != 999 {
		t.Fatalf("expected 999, got %v", updated.Reserved)
	}
}

func TestTransactionRepository_Approve(t *testing.T) {
	tr := domain.Transaction{
		ID:        "1",
		Reserved:  600,
		ProductID: 1,
		VendorID:  1,
		Status:    domain.TRANSACTION_DRAFT,
	}

	if err := transactionRepo.Create(tr); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer transactionRepo.DeleteTransactionsByIDs(tr.ID)

	if err := transactionRepo.Approve(tr.ID); err != nil {
		t.Fatalf("approve failed: %v", err)
	}

	updated, err := transactionRepo.GetByID(tr.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if updated.Status != domain.TRANSACTION_SUCCESS {
		t.Fatalf("expected SUCCESS, got %v", updated.Status)
	}
}

func TestTransactionRepository_Reject(t *testing.T) {
	tr := domain.Transaction{
		ID:        "1",
		Reserved:  600,
		ProductID: 1,
		VendorID:  1,
		Status:    domain.TRANSACTION_DRAFT,
	}

	if err := transactionRepo.Create(tr); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer transactionRepo.DeleteTransactionsByIDs(tr.ID)

	if err := transactionRepo.Reject(tr.ID); err != nil {
		t.Fatalf("approve failed: %v", err)
	}

	updated, err := transactionRepo.GetByID(tr.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if updated.Status != domain.TRANSACTION_FAIL {
		t.Fatalf("expected SUCCESS, got %v", updated.Status)
	}
}
