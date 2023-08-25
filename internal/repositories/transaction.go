package repositories

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transactionRepository {
	return transactionRepository{db}
}

func (r transactionRepository) BeginTx() *gorm.DB {
	return r.db.Begin()
}

func (r transactionRepository) CommitTx(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r transactionRepository) RollbackTx(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r transactionRepository) Create(tx *gorm.DB, newTransaction *models.Transaction) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	newTransaction.ID = uuid.New()
	newTransaction.Status = enum.TransactionStatusSuccess
	newTransaction.TransactedAt = time.Now()

	return db.Model(newTransaction).Create(newTransaction).Error
}

func (r transactionRepository) ListByWalletID(walletID uuid.UUID) ([]models.Transaction, error) {
	var transactions []models.Transaction

	db := r.db.Model(&transactions)
	if err := db.Where("wallet_id = ?", walletID.String()).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
