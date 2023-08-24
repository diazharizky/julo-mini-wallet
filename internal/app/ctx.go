package app

type Ctx struct {
	AccountRepository     IAccountRepository
	WalletRepository      IWalletRepository
	TransactionRepository ITransactionRepository

	InitializeAccountModule       IInitializeAccountModule
	EnableWalletModule            IEnableWalletModule
	ListWalletTransactionsModule  IListWalletTransactionsModule
	GenerateTokenModule           IGenerateTokenModule
	ValidateTokenModule           IValidateTokenModule
	DepositWalletBalanceModule    IDepositWalletBalanceModule
	WithdrawalWalletBalanceModule IWithdrawalWalletBalanceModule
	DisableWalletModule           IDisableWalletModule
}
