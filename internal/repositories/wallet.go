package repositories

import (
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type walletRepository struct{}

func NewWalletRepository() walletRepository {
	return walletRepository{}
}

func (walletRepository) GetByAccountID(accountID uuid.UUID) (wallet *models.Wallet, err error) {
	return nil, nil
}

func (walletRepository) Create(accountID uuid.UUID) (newWallet *models.Wallet, err error) {
	newWallet = models.NewDefaultWAllet(accountID)

	return
}
