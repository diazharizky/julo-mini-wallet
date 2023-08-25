package repositories

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) accountRepository {
	return accountRepository{db}
}

func (r accountRepository) BeginTx() *gorm.DB {
	return r.db.Begin()
}

func (r accountRepository) CommitTx(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r accountRepository) RollbackTx(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r accountRepository) Create(tx *gorm.DB, newAccount *models.Account) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	newAccount.CreatedAt = time.Now()

	return db.Model(newAccount).Create(newAccount).Error
}
