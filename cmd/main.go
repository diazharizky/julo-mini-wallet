package main

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/modules"
	"github.com/diazharizky/julo-mini-wallet/internal/repositories"
	"github.com/diazharizky/julo-mini-wallet/internal/server"
)

func main() {
	appCtx := app.Ctx{}

	appCtx.UserRepository = repositories.NewUserRepository()
	appCtx.WalletRepository = repositories.NewWalletRepository()
	appCtx.TransactionRepository = repositories.NewTransactionRepository()

	appCtx.InitializeAccountModule = modules.NewInitializeAccountModule(appCtx)
	appCtx.EnableWalletModule = modules.NewEnableWalletModule(appCtx)
	appCtx.ListWalletTransactionsModule = modules.NewListWalletTransactionsModule(appCtx)
	appCtx.GenerateTokenModule = modules.NewGenerateTokenModule()
	appCtx.ValidateTokenModule = modules.NewValidateTokenModule()
	appCtx.DepositWalletBalanceModule = modules.NewDepositWalletBalanceModule(appCtx)

	svr := server.New(appCtx)

	svr.Listen(":80")
}
