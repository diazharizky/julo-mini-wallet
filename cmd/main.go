package main

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/modules"
	"github.com/diazharizky/julo-mini-wallet/internal/repositories"
	"github.com/diazharizky/julo-mini-wallet/internal/server"
	"github.com/diazharizky/julo-mini-wallet/pkg/db"
)

func main() {
	appCtx := app.Ctx{}

	db := db.New().DB()

	appCtx.AccountRepository = repositories.NewAccountRepository(db)
	appCtx.WalletRepository = repositories.NewWalletRepository(db)
	appCtx.TransactionRepository = repositories.NewTransactionRepository(db)

	appCtx.InitAccountModule = modules.NewInitAccountModule(appCtx)
	appCtx.EnableWalletModule = modules.NewEnableWalletModule(appCtx)
	appCtx.ListWalletTransactionsModule = modules.NewListWalletTransactionsModule(appCtx)
	appCtx.DepositWalletBalanceModule = modules.NewDepositWalletBalanceModule(appCtx)
	appCtx.WithdrawalWalletBalanceModule = modules.NewWithdrawalWalletBalanceModule(appCtx)
	appCtx.DisableWalletModule = modules.NewDisableWalletModule(appCtx)

	svr := server.New(appCtx)

	svr.Listen(":80")
}
