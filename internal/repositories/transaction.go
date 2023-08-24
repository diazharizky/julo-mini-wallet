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

func (transactionRepository) List(walletID uuid.UUID) ([]models.Transaction, error) {
	transactions := []models.Transaction{
		{
			ID:           uuid.New(),
			WalletID:     walletID,
			Status:       enum.TransactionStatusSuccess,
			TransactedAt: time.Now(),
			Type:         enum.TransactionTypeDeposit,
			Amount:       14000,
			ReferenceID:  uuid.New(),
		},
		{
			ID:           uuid.New(),
			WalletID:     walletID,
			Status:       enum.TransactionStatusFailed,
			TransactedAt: time.Now(),
			Type:         enum.TransactionTypeWithdrawal,
			Amount:       890000,
			ReferenceID:  uuid.New(),
		},
	}

	return transactions, nil
}
