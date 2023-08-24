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

	db := r.db.Model(&models.Wallet{})
	if err := db.First(&wallet, "owned_by = ?", accountID.String()).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if wallet.ID == uuid.Nil {
		return nil, nil
	}

	return &wallet, nil
}

func (r walletRepository) Create(accountID uuid.UUID) (*models.Wallet, error) {
	newWallet := models.NewDefaultWAllet(accountID)

	db := r.db.Model(&models.Wallet{})
	if err := db.Create(&newWallet).Error; err != nil {
		return nil, err
	}

	return newWallet, nil
}

func (r walletRepository) Enable(wallet *models.Wallet) error {
	now := time.Now()

	wallet.DisabledAt = nil

	db := r.db.Model(wallet)
	return db.
		Where("id = ?", wallet.ID.String()).
		Updates(map[string]interface{}{
			"status":      enum.WalletStatusEnabled,
			"enabled_at":  &now,
			"disabled_at": gorm.Expr("NULL"),
		}).
		Error
}

func (r walletRepository) Disable(wallet *models.Wallet) error {
	now := time.Now()

	wallet.EnabledAt = nil

	db := r.db.Model(wallet)
	return db.
		Where("id = ?", wallet.ID.String()).
		Updates(map[string]interface{}{
			"status":      enum.WalletStatusDisabled,
			"disabled_at": &now,
			"enabled_at":  gorm.Expr("NULL"),
		}).
		Error
}

func (r walletRepository) Deposit(wallet *models.Wallet, depositAmount float64) error {
	db := r.db.Model(wallet)
	return db.Where("id = ?", wallet.ID.String()).Update("balance", wallet.Balance+depositAmount).Error
}
