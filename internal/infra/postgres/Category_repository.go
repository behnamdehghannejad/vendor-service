package postgres

import (
	"strconv"
	"strings"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repo *CategoryRepository) Create(category domain.Category) (int, error) {
	path := ""
	if category.ParentID != nil {
		parent, _ := repo.FindById(*category.ParentID)
		path = parent.Path
	}

	entity := toCategoryModel(category, path)

	if err := repo.db.Create(entity).Error; err != nil {
		return 0, convertPostgresErrorToAppError(err, category)
	}

	return entity.ID, nil
}

func (repo *CategoryRepository) Update(category domain.Category) error {
	if err := repo.db.Save(toCategoryModel(category, "")).Error; err != nil {
		return convertPostgresErrorToAppError(err, category)
	}

	return nil
}

func (repo *CategoryRepository) SoftDelete(id int) error {
	if err := repo.db.Model(&model.CategoryModel{}).
		Where("id = ?", id).
		Update("active", false).
		Error; err != nil {
		return convertPostgresErrorToAppError(err, id)
	}

	return nil
}

func (repo *CategoryRepository) FindById(id int) (domain.Category, error) {
	var entity model.CategoryModel

	if err := repo.db.First(&entity, id).Error; err != nil {
		return domain.Category{}, convertPostgresErrorToAppError(err)
	}

	return toCategoryDomain(entity), nil
}

func (repo *CategoryRepository) FindChildren(id int) ([]domain.Category, error) {
	var parent model.CategoryModel

	if err := repo.db.First(&parent, id).Error; err != nil {
		return nil, convertPostgresErrorToAppError(err, id)
	}

	var entities []model.CategoryModel

	if err := repo.db.
		Where("path LIKE ?", parent.Path+"/%").
		Where("id <> ?", parent.ID).
		Find(&entities).
		Error; err != nil {
		return nil, convertPostgresErrorToAppError(err, id)
	}

	result := make([]domain.Category, 0, len(entities))
	for _, entity := range entities {
		result = append(result, toCategoryDomain(entity))
	}

	return result, nil
}

func (repo *CategoryRepository) FindParents(id int) ([]domain.Category, error) {
	category, err := repo.FindById(id)
	if err != nil {
		return nil, convertPostgresErrorToAppError(err, id)
	}

	ids := parsePath(category.Path)

	var entities []model.CategoryModel

	if err := repo.db.
		Where("id IN ?", ids).
		Order("id").
		Find(&entities).
		Error; err != nil {
		return nil, convertPostgresErrorToAppError(err, id)
	}

	result := make([]domain.Category, 0, len(entities))
	for _, entity := range entities {
		result = append(result, toCategoryDomain(entity))
	}

	return result, nil
}

func toCategoryModel(category domain.Category, path string) *model.CategoryModel {
	return &model.CategoryModel{
		ID:       category.ID,
		Name:     category.Name,
		Active:   category.Active,
		ParentID: category.ParentID,
		Path:     getPath(category, path),
	}
}

func toCategoryDomain(category model.CategoryModel) domain.Category {
	return domain.Category{
		ID:       category.ID,
		Name:     category.Name,
		Active:   category.Active,
		ParentID: category.ParentID,
		Path:     category.Path,
	}
}

func getPath(category domain.Category, parentsPath string) string {
	path := ""
	if category.Path != "" {
		path = category.Path
	} else if parentsPath != "" {
		path = parentsPath + "/" + strconv.Itoa(*category.ParentID)
	} else if category.ParentID != nil {
		path = strconv.Itoa(*category.ParentID)
	}
	return path
}

func parsePath(path string) []int {
	parts := strings.Split(strings.Trim(path, "/"), "/")

	result := make([]int, 0, len(parts))

	for _, part := range parts {
		id, err := strconv.Atoi(part)
		if err == nil {
			result = append(result, id)
		}
	}

	return result
}

func (repo *CategoryRepository) Delete(id int) error {
	if err := repo.db.Delete(&model.CategoryModel{}, id).Error; err != nil {
		return convertPostgresErrorToAppError(err, id)
	}
	return nil
}
