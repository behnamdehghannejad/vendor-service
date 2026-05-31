package postgres

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres/model"

	"gorm.io/gorm"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{
		db: db,
	}
}

func (repo *HistoryRepository) Create(history domain.History) error {
	if err := repo.db.Create(repo.toHistoryEntity(history)).Error; err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *HistoryRepository) Update(history domain.History) error {
	if err := repo.db.Save(repo.toHistoryEntity(history)).Error; err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *HistoryRepository) FindByOrderID(id string) (domain.History, error) {
	var history model.HistoryModel
	if err := repo.db.Where("order_id = ?", id).First(&history).Error; err != nil {
		return domain.History{}, convertPostgresErrorToAppError(err)
	}
	domainHistory := repo.toHistoryDomain(history)
	return domainHistory, nil
}

func (repo *HistoryRepository) Filter(filter domain.SearchHistory) ([]domain.History, error) {
	q := repo.db.Model(&model.HistoryModel{})

	if filter.Activation != nil {
		q = q.Where("active = ?", *filter.Activation)
	}

	if filter.PaymentID != "" {
		q = q.Where("payment_id = ?", filter.PaymentID)
	}

	if filter.OrderID != "" {
		q = q.Where("order_id = ?", filter.OrderID)
	}

	if filter.VendorID != nil {
		q = q.Where("vendor_id = ?", *filter.VendorID)
	}

	if filter.ProductID != nil {
		q = q.Where("product_id = ?", *filter.ProductID)
	}

	if filter.Status != nil {
		q = q.Where("status = ?", *filter.Status)
	}

	var histories []model.HistoryModel

	err := q.Find(&histories).Error
	if err != nil {
		return nil, convertPostgresErrorToAppError(err)
	}

	return repo.toHistoryDomains(histories), err
}

func (repo *HistoryRepository) toHistoryDomains(historyModels []model.HistoryModel) []domain.History {
	histories := make([]domain.History, 0, len(historyModels))
	for _, history := range histories {
		histories = append(histories, history)
	}
	return histories
}

func (repo *HistoryRepository) toHistoryEntity(domain domain.History) model.HistoryModel {
	return model.HistoryModel{
		Reserved:  domain.Reserved,
		ProductID: domain.ProductID,
		VendorID:  domain.VendorID,
	}
}

func (repo *HistoryRepository) toHistoryDomain(history model.HistoryModel) domain.History {
	return domain.History{
		ID:        history.ID,
		Reserved:  history.Reserved,
		ProductID: history.ProductID,
		VendorID:  history.VendorID,
		Status:    history.Status,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
	}
}
