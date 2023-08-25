package app

import (
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAccountRepository interface {
	BeginTx() *gorm.DB
	CommitTx(tx *gorm.DB) error
	RollbackTx(tx *gorm.DB) error
	Create(tx *gorm.DB, newAccount *models.Account) error
}

type IWalletRepository interface {
	BeginTx() *gorm.DB
	CommitTx(tx *gorm.DB) error
	RollbackTx(tx *gorm.DB) error
	GetByAccountID(accountID uuid.UUID, clauses ...clause.Expression) (wallet *models.Wallet, err error)
	Create(tx *gorm.DB, accountID uuid.UUID) (newWallet *models.Wallet, err error)
	Enable(tx *gorm.DB, wallet *models.Wallet) error
	Disable(tx *gorm.DB, wallet *models.Wallet) error
	DepositBalance(tx *gorm.DB, wallet *models.Wallet, depositAmount float64) error
}

type ITransactionRepository interface {
	BeginTx() *gorm.DB
	CommitTx(tx *gorm.DB) error
	RollbackTx(tx *gorm.DB) error
	Create(tx *gorm.DB, newTransaction *models.Transaction) error
	ListByWalletID(walletID uuid.UUID) (transactions []models.Transaction, err error)
}

type IInitAccountModule interface {
	Call(id uuid.UUID) error
}

type IEnableWalletModule interface {
	Call(accountID uuid.UUID) (newWallet *models.Wallet, err error)
}

type IListWalletTransactionsModule interface {
	Call(accountID uuid.UUID) (transactions []models.Transaction, err error)
}

type IDepositWalletBalanceModule interface {
	Call(newDeposit *models.Deposit) error
}

type IWithdrawalWalletBalanceModule interface {
	Call(newWithdrawal *models.Withdrawal) error
}

type IDisableWalletModule interface {
	Call(accountID uuid.UUID) (*models.Wallet, error)
}
