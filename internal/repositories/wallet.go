package repositories

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type walletRepository struct{}

func NewWalletRepository() walletRepository {
	return walletRepository{}
}

func (walletRepository) GetByAccountID(accountID uuid.UUID) (wallet *models.Wallet, err error) {
	wallet = models.NewDefaultWAllet(accountID)

	return
}

func (walletRepository) Create(accountID uuid.UUID) (newWallet *models.Wallet, err error) {
	newWallet = models.NewDefaultWAllet(accountID)

	return
}

func (walletRepository) Disable(wallet *models.Wallet) error {
	now := time.Now()
	wallet.Status = enum.WalletStatusDisabled
	wallet.DisabledAt = &now

	return nil
}
