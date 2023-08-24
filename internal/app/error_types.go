package app

const (
	WalletIsAlreadyEnabled = err("wallet is already enabled")
	InsufficientBalance    = err("insufficient balance")
	WalletIsDisabled       = err("wallet is disabled")
)
