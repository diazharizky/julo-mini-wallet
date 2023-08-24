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

func (r transactionRepository) Create(newTransaction *models.Transaction) error {
	newTransaction.ID = uuid.New()
	newTransaction.Status = enum.TransactionStatusSuccess
	newTransaction.TransactedAt = time.Now()

	return r.db.Create(&newTransaction).Error
}

func (r transactionRepository) List(walletID uuid.UUID) ([]models.Transaction, error) {
	var transactions []models.Transaction

	db := r.db.Model(&transactions)
	if err := db.Where("wallet_id = ?", walletID.String()).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
