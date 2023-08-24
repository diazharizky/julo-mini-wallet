package repositories

import (
	"errors"
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) walletRepository {
	return walletRepository{db}
}

func (r walletRepository) GetByAccountID(accountID uuid.UUID) (*models.Wallet, error) {
	var wallet models.Wallet

	if err := r.db.First(&wallet, "owned_by = ?", accountID.String()).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if wallet.ID == uuid.Nil {
		return nil, nil
	}

	return &wallet, nil
}

func (r walletRepository) Create(accountID uuid.UUID) (*models.Wallet, error) {
	newWallet := models.NewDefaultWAllet(accountID)

	if err := r.db.Create(&newWallet).Error; err != nil {
		return nil, err
	}

	return newWallet, nil
}

func (walletRepository) Disable(wallet *models.Wallet) error {
	now := time.Now()
	wallet.Status = enum.WalletStatusDisabled
	wallet.DisabledAt = &now

	return nil
}
