package app

import (
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type IUserRepository interface{}

type IWalletRepository interface {
	GetByAccountID(accountID uuid.UUID) (wallet *models.Wallet, err error)
	Create(accountID uuid.UUID) (newWallet *models.Wallet, err error)
}

type IInitializeAccountModule interface {
	Call(customerXID uuid.UUID)
}

type IEnableWalletModule interface {
	Call(accountID uuid.UUID) (newWallet *models.Wallet, err error)
}

type IGenerateTokenModule interface {
	Call(accountId uuid.UUID) string
}

type IValidateTokenModule interface {
	Call(token string) (string, error)
}
