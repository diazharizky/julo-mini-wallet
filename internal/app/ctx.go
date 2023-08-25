package app

type Ctx struct {
	AccountRepository     IAccountRepository
	WalletRepository      IWalletRepository
	TransactionRepository ITransactionRepository

	InitAccountModule             IInitAccountModule
	EnableWalletModule            IEnableWalletModule
	ListWalletTransactionsModule  IListWalletTransactionsModule
	DepositWalletBalanceModule    IDepositWalletBalanceModule
	WithdrawalWalletBalanceModule IWithdrawalWalletBalanceModule
	DisableWalletModule           IDisableWalletModule
}
