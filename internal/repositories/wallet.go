package repositories

import (
	"errors"
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) walletRepository {
	return walletRepository{db}
}

func (r walletRepository) BeginTx() *gorm.DB {
	return r.db.Begin()
}

func (r walletRepository) CommitTx(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r walletRepository) RollbackTx(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r walletRepository) GetByAccountID(accountID uuid.UUID, clauses ...clause.Expression) (*models.Wallet, error) {
	var wallet models.Wallet

	if len(clauses) > 0 {
		r.db = r.db.Clauses(clauses...)
	}

	if err := r.db.Model(wallet).First(&wallet, "owned_by = ?", accountID.String()).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if wallet.ID == uuid.Nil {
		return nil, nil
	}

	return &wallet, nil
}

func (r walletRepository) Create(tx *gorm.DB, accountID uuid.UUID) (*models.Wallet, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	newWallet := models.NewDefaultWAllet(accountID)

	if err := db.Model(newWallet).Create(newWallet).Error; err != nil {
		return nil, err
	}

	return newWallet, nil
}

func (r walletRepository) Enable(tx *gorm.DB, wallet *models.Wallet) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	now := time.Now()

	wallet.DisabledAt = nil

	return db.
		Model(wallet).
		Where("id = ?", wallet.ID.String()).
		Updates(map[string]interface{}{
			"status":      enum.WalletStatusEnabled,
			"enabled_at":  &now,
			"disabled_at": gorm.Expr("NULL"),
		}).
		Error
}

func (r walletRepository) Disable(tx *gorm.DB, wallet *models.Wallet) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	now := time.Now()

	wallet.EnabledAt = nil

	return db.
		Model(wallet).
		Where("id = ?", wallet.ID.String()).
		Updates(map[string]interface{}{
			"status":      enum.WalletStatusDisabled,
			"disabled_at": &now,
			"enabled_at":  gorm.Expr("NULL"),
		}).
		Error
}

func (r walletRepository) DepositBalance(tx *gorm.DB, wallet *models.Wallet, depositAmount float64) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	return db.
		Model(wallet).
		Where("id = ?", wallet.ID.String()).
		Update("balance", wallet.Balance+depositAmount).
		Error
}
