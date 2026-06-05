package repository

import (
	"testing"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
)

func TestCategoryRepository_Create(t *testing.T) {
	category := domain.Category{
		Name:   "Electronics",
		Active: true,
	}

	id, err := categoryRepo.Create(category)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	categoryRepo.Delete(id)
}

func TestCategoryRepository_FindById(t *testing.T) {
	category := domain.Category{
		Name:   "Mobiles",
		Active: true,
	}

	id, err := categoryRepo.Create(category)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer categoryRepo.Delete(id)

	found, err := categoryRepo.FindById(id)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if found.Name != "Mobiles" {
		t.Fatalf("expected name 'Mobiles', got %s", found.Name)
	}
}

func TestCategoryRepository_Update(t *testing.T) {
	category := domain.Category{
		Name:   "Old Category",
		Active: true,
	}

	id, err := categoryRepo.Create(category)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer categoryRepo.Delete(id)

	category.ID = id
	category.Name = "New Category"

	if err := categoryRepo.Update(category); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	updated, err := categoryRepo.FindById(id)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}

	if updated.Name != "New Category" {
		t.Fatalf("expected updated name, got %s", updated.Name)
	}
}

func TestCategoryRepository_SoftDelete(t *testing.T) {
	category := domain.Category{
		Name:   "Delete Me",
		Active: true,
	}

	id, err := categoryRepo.Create(category)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	defer categoryRepo.Delete(id)

	if err := categoryRepo.SoftDelete(id); err != nil {
		t.Fatalf("soft delete failed: %v", err)
	}

	found, err := categoryRepo.FindById(id)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}

	if found.Active {
		t.Fatal("expected category to be inactive")
	}
}

func TestCategoryRepository_FindChildren(t *testing.T) {
	parentID, err := categoryRepo.Create(domain.Category{
		Name:   "Parent",
		Active: true,
	})
	if err != nil {
		t.Fatalf("create parent failed: %v", err)
	}

	childID, err := categoryRepo.Create(domain.Category{
		Name:     "Child",
		ParentID: &parentID,
		Active:   true,
	})
	if err != nil {
		t.Fatalf("create child failed: %v", err)
	}

	defer categoryRepo.Delete(childID)
	defer categoryRepo.Delete(parentID)

	children, err := categoryRepo.FindChildren(parentID)
	if err != nil {
		t.Fatalf("find children failed: %v", err)
	}

	if len(children) != 1 {
		t.Fatalf("expected 1 child, got %d", len(children))
	}

	if children[0].ID != childID {
		t.Fatalf("expected child id %d, got %d", childID, children[0].ID)
	}
}

func TestCategoryRepository_FindParents(t *testing.T) {
	parentID, err := categoryRepo.Create(domain.Category{
		Name:   "Parent",
		Active: true,
	})
	if err != nil {
		t.Fatalf("create parent failed: %v", err)
	}

	childID, err := categoryRepo.Create(domain.Category{
		Name:     "Child",
		ParentID: &parentID,
		Active:   true,
	})
	if err != nil {
		t.Fatalf("create child failed: %v", err)
	}

	defer categoryRepo.Delete(childID)
	defer categoryRepo.Delete(parentID)

	parents, err := categoryRepo.FindParents(childID)
	if err != nil {
		t.Fatalf("find parents failed: %v", err)
	}

	if len(parents) != 1 {
		t.Fatalf("expected 1 parent, got %d", len(parents))
	}

	if parents[0].ID != parentID {
		t.Fatalf("expected parent id %d, got %d", parentID, parents[0].ID)
	}
}
