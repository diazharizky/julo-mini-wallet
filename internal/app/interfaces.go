package app

import (
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type IUserRepository interface{}

type IWalletRepository interface {
	GetByAccountID(accountID uuid.UUID) (wallet *models.Wallet, err error)
	Create(accountID uuid.UUID) (newWallet *models.Wallet, err error)
	Disable(wallet *models.Wallet) error
}

type ITransactionRepository interface {
	List(walletID uuid.UUID) (transactions []models.Transaction, err error)
}

type IInitializeAccountModule interface {
	Call(customerXID uuid.UUID)
}

type IEnableWalletModule interface {
	Call(accountID uuid.UUID) (newWallet *models.Wallet, err error)
}

type IListWalletTransactionsModule interface {
	Call(accountID uuid.UUID) (transactions []models.Transaction, err error)
}

type IGenerateTokenModule interface {
	Call(accountId uuid.UUID) string
}

type IValidateTokenModule interface {
	Call(token string) (string, error)
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
