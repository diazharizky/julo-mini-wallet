package app

type Ctx struct {
	WalletRepository      IWalletRepository
	TransactionRepository ITransactionRepository

	EnableWalletModule            IEnableWalletModule
	ListWalletTransactionsModule  IListWalletTransactionsModule
	DepositWalletBalanceModule    IDepositWalletBalanceModule
	WithdrawalWalletBalanceModule IWithdrawalWalletBalanceModule
	DisableWalletModule           IDisableWalletModule
}
