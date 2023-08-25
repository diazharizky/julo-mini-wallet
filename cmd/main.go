package main

import (
	"fmt"

	"github.com/diazharizky/julo-mini-wallet/config"
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/modules"
	"github.com/diazharizky/julo-mini-wallet/internal/repositories"
	"github.com/diazharizky/julo-mini-wallet/internal/server"
	"github.com/diazharizky/julo-mini-wallet/pkg/db"
)

func init() {
	config.Global.SetDefault("server.host", "localhost")
	config.Global.SetDefault("server.port", "8080")
}

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

	addr := fmt.Sprintf(
		"%s:%s",
		config.Global.GetString("server.host"),
		config.Global.GetString("server.port"),
	)

	svr.Listen(addr)
}
