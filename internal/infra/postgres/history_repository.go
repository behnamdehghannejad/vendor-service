package postgres

import (
	"time"

	"github.com/behnamdehghannejad/vendorservice/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HistoryEntity struct {
	ID        int           `gorm:"primary_key"`
	OrderID   uuid.UUID     `gorm:"column:order_id"`
	PaymentID uuid.UUID     `gorm:"column:payment_id"`
	Quantity  int           `gorm:"column:quantity"`
	ProductID int           `gorm:"column:product_id"`
	VendorID  int           `gorm:"column:vendor_id"`
	Status    domain.Status `gorm:"column:status"`
	Active    bool          `gorm:"column:active"`
	CreatedAt time.Time     `gorm:column"created_at"`
	UpdatedAt time.Time     `gorm:column"updated_at"`
}

func (HistoryEntity) HistoryTableName() string {
	return "history"
}

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{
		db: db,
	}
}

func (repo *HistoryRepository) Add(history domain.History) error {
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

func (repo *HistoryRepository) Delete(id int) error {
	err := repo.db.Model(&HistoryEntity{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"active":     false,
		}).Error
	if err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *HistoryRepository) FindByOrderID(id uuid.UUID) (domain.History, error) {
	var history HistoryEntity
	if err := repo.db.Where("order_id = ?", id).First(&history).Error; err != nil {
		return domain.History{}, convertPostgresErrorToAppError(err)
	}
	domainHistory := repo.toHistoryDomain(history)
	return domainHistory, nil
}

func (repo *HistoryRepository) FindByPaymentID(paymentID uuid.UUID) (domain.History, error) {
	var history HistoryEntity
	if err := repo.db.Where("payment_id = ?", paymentID).First(&history).Error; err != nil {
		return domain.History{}, convertPostgresErrorToAppError(err)
	}
	domainHistory := repo.toHistoryDomain(history)
	return domainHistory, nil
}

func (repo *HistoryRepository) FindByProductID(productID int) ([]domain.History, error) {
	var historyList []HistoryEntity
	if err := repo.db.Where("product_id = ?", productID).Find(&historyList).Error; err != nil {
		return nil, convertPostgresErrorToAppError(err)
	}

	domainHistoryList := repo.addToDomainList(historyList)
	return domainHistoryList, nil
}

func (repo *HistoryRepository) FindByVendorID(vendorID int) ([]domain.History, error) {
	var historyList []HistoryEntity
	if err := repo.db.Where("vendor_id = ?", vendorID).Find(&historyList).Error; err != nil {
		return nil, convertPostgresErrorToAppError(err)
	}

	domainHistoryList := repo.addToDomainList(historyList)
	return domainHistoryList, nil
}

func (repo *HistoryRepository) FindByStatus(status domain.Status) ([]domain.History, error) {
	var historyList []HistoryEntity
	if err := repo.db.Where("status = ?", status).Find(&historyList).Error; err != nil {
		return nil, convertPostgresErrorToAppError(err)
	}

	domainHistoryList := repo.addToDomainList(historyList)
	return domainHistoryList, nil
}

func (repo *HistoryRepository) FindByIsActive(isActive bool) ([]domain.History, error) {
	var historyList []HistoryEntity
	if err := repo.db.Where("is_active = ?", isActive).Find(&historyList).Error; err != nil {
		return nil, convertPostgresErrorToAppError(err)
	}

	domainHistoryList := repo.addToDomainList(historyList)
	return domainHistoryList, nil
}

func (repo *HistoryRepository) addToDomainList(historyList []HistoryEntity) []domain.History {
	domainHistoryList := make([]domain.History, 0, len(historyList))
	for _, history := range historyList {
		domainHistoryList = append(domainHistoryList, repo.toHistoryDomain(history))
	}
	return domainHistoryList
}

func (repo *HistoryRepository) toHistoryEntity(domain domain.History) HistoryEntity {
	return HistoryEntity{
		OrderID:   domain.OrderID,
		PaymentID: domain.PaymentID,
		Quantity:  domain.Quantity,
		ProductID: domain.ProductID,
		VendorID:  domain.VendorID,
		Active:    domain.Active,
		CreatedAt: time.Now(),
	}
}

func (repo *HistoryRepository) toHistoryDomain(history HistoryEntity) domain.History {
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
