package postgres

import (
	"github.com/behnamdehghannejad/vendorservice/internal/domain"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres/model"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (repo *TransactionRepository) Create(transaction domain.Transaction) error {
	transactionModel := repo.toHistoryModel(transaction)
	if err := repo.db.Create(&transactionModel).Error; err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *TransactionRepository) Update(transaction domain.Transaction) error {
	if err := repo.db.Save(repo.toHistoryModel(transaction)).Error; err != nil {
		return convertPostgresErrorToAppError(err)
	}
	return nil
}

func (repo *TransactionRepository) GetByID(ID string) (domain.Transaction, error) {
	var transactionModel model.TransactionModel
	if err := repo.db.Where("id = ?", ID).First(&transactionModel).Error; err != nil {
		return domain.Transaction{}, convertPostgresErrorToAppError(err)
	}

	return repo.toHistoryDomain(transactionModel), nil
}

func (repo *TransactionRepository) Filter(filter domain.SearchTransaction) ([]domain.Transaction, error) {
	q := repo.db.Model(&model.TransactionModel{})

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

	var transactions []model.TransactionModel

	err := q.Find(&transactions).Error
	if err != nil {
		return nil, convertPostgresErrorToAppError(err)
	}

	return repo.toHistoryDomains(transactions), err
}

func (repo *TransactionRepository) Approve(ID string) error {
	err := repo.db.Model(&model.TransactionModel{}).
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"status":     domain.HISTORY_SUCCESS,
			"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
		}).Error
	if err != nil {
		return apperror.Wrap(err).
			UnExpected().
			Log().
			Build()
	}
	return nil
}

func (repo *TransactionRepository) Reject(ID string) error {
	err := repo.db.Model(&model.TransactionModel{}).
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"status":     domain.HISTORY_FAIL,
			"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
		}).Error
	if err != nil {
		return apperror.Wrap(err).
			UnExpected().
			Log().
			Build()
	}
	return nil
}

func (repo *TransactionRepository) toHistoryDomains(transactionModels []model.TransactionModel) []domain.Transaction {
	transactions := make([]domain.Transaction, 0, len(transactionModels))
	for _, transactionModel := range transactionModels {
		transactions = append(transactions, repo.toHistoryDomain(transactionModel))
	}
	return transactions
}

func (repo *TransactionRepository) toHistoryModel(domain domain.Transaction) model.TransactionModel {
	return model.TransactionModel{
		ID:        domain.ID,
		Reserved:  domain.Reserved,
		ProductID: domain.ProductID,
		VendorID:  domain.VendorID,
	}
}

func (repo *TransactionRepository) toHistoryDomain(transaction model.TransactionModel) domain.Transaction {
	return domain.Transaction{
		ID:        transaction.ID,
		Reserved:  transaction.Reserved,
		ProductID: transaction.ProductID,
		VendorID:  transaction.VendorID,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
	}
}
