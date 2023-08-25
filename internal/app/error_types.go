package app

const (
	WalletIsAlreadyEnabled = err("app: wallet is already enabled")
	InsufficientBalance    = err("app: insufficient balance")
	WalletIsDisabled       = err("app: wallet is disabled")
)
