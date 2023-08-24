package app

import "github.com/google/uuid"

type IUserRepository interface{}

type IWalletRepository interface{}

type IInitializeAccountModule interface {
	Call(customerXID uuid.UUID)
}

type IEnableWalletModule interface {
	Call()
}

type IGenerateTokenModule interface {
	Call(accountId uuid.UUID) string
}

type IValidateTokenModule interface {
	Call(token string) (string, error)
}
