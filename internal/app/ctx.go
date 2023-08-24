package app

type Ctx struct {
	UserRepository   IUserRepository
	WalletRepository IWalletRepository

	InitializeAccountModule IInitializeAccountModule
	EnableWalletModule      IEnableWalletModule
	GenerateTokenModule     IGenerateTokenModule
	ValidateTokenModule     IValidateTokenModule
}
