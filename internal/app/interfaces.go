package app

import (
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type IWalletRepository interface {
	GetByAccountID(accountID uuid.UUID) (wallet *models.Wallet, err error)
	Create(accountID uuid.UUID) (newWallet *models.Wallet, err error)
	Enable(wallet *models.Wallet) error
	Disable(wallet *models.Wallet) error
	Deposit(wallet *models.Wallet, depositAmount float64) error
}

type ITransactionRepository interface {
	Create(newTransaction *models.Transaction) error
	List(walletID uuid.UUID) (transactions []models.Transaction, err error)
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
