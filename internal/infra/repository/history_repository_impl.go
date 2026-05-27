package repository

import (
	"time"

	"github.com/behnamdehghannejad/vendor/internal/domain"

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

type HistoryRepositoryImpl struct {
	db *gorm.DB
}

func NewHistoryRepositoryImpl(db *gorm.DB) *HistoryRepositoryImpl {
	return &HistoryRepositoryImpl{
		db: db,
	}
}

func (repo *HistoryRepositoryImpl) Add(history *domain.History) error {
	return repo.db.Create(toHistoryEntity(history)).Error
}

func (repo *HistoryRepositoryImpl) Update(history *domain.History) error {
	return repo.db.Save(toHistoryEntity(history)).Error
}

func (repo *HistoryRepositoryImpl) Delete(id int) error {
	var history HistoryEntity
	if err := repo.db.Where("id = ?", id).Find(&history).Error; err != nil {
		return err
	}

	history.UpdatedAt = time.Now()
	history.Active = false
	return repo.db.Save(&history).Error
}

func (repo *HistoryRepositoryImpl) FindByOrderID(id uuid.UUID) (*domain.History, error) {
	var history HistoryEntity
	if err := repo.db.Where("order_id = ?", id).First(&history).Error; err != nil {
		return nil, err
	}
	return toHistoryDomain(&history), nil
}

func (repo *HistoryRepositoryImpl) FindByPaymentID(paymentID uuid.UUID) (*domain.History, error) {
	var history HistoryEntity
	if err := repo.db.Where("payment_id = ?", paymentID).First(&history).Error; err != nil {
		return nil, err
	}
	return toHistoryDomain(&history), nil
}

func (repo *HistoryRepositoryImpl) FindByProductID(productID int) ([]domain.History, error) {
	var historyList []HistoryEntity
	if err := repo.db.Where("product_id = ?", productID).Find(&historyList).Error; err != nil {
		return nil, err
	}

	domainHistoryList := repo.addToDomainList(historyList)
	return domainHistoryList, nil
}

func (repo *HistoryRepositoryImpl) FindByVendorID(vendorID int) ([]domain.History, error) {
	var historyList []HistoryEntity
	if err := repo.db.Where("vendor_id = ?", vendorID).Find(&historyList).Error; err != nil {
		return nil, err
	}

	domainHistoryList := repo.addToDomainList(historyList)
	return domainHistoryList, nil
}

func (repo *HistoryRepositoryImpl) FindByStatus(status domain.Status) ([]domain.History, error) {
	var historyList []HistoryEntity
	if err := repo.db.Where("status = ?", status).Find(&historyList).Error; err != nil {
		return nil, err
	}

	domainHistoryList := repo.addToDomainList(historyList)
	return domainHistoryList, nil
}

func (repo *HistoryRepositoryImpl) FindByIsActive(isActive bool) ([]domain.History, error) {
	var historyList []HistoryEntity
	if err := repo.db.Where("is_active = ?", isActive).Find(&historyList).Error; err != nil {
		return nil, err
	}

	domainHistoryList := repo.addToDomainList(historyList)
	return domainHistoryList, nil
}

func (repo *HistoryRepositoryImpl) addToDomainList(historyList []HistoryEntity) []domain.History {
	domainHistoryList := make([]domain.History, 0, len(historyList))
	for _, history := range historyList {
		domainHistoryList = append(domainHistoryList, toHistoryDomainValue(history))
	}
	return domainHistoryList
}
