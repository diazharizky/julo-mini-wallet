package server

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/server/controllers"
	"github.com/diazharizky/julo-mini-wallet/internal/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func New(appCtx app.Ctx) (svr *fiber.App) {
	svr = fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	api := svr.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/init", controllers.InitAccountController(appCtx))

	v1Protected := v1.Use(middlewares.AuthorizationMiddleware(appCtx))

	walletEndpoint := v1Protected.Group("/wallet")
	{
		walletEndpoint.Post("/", controllers.EnableWalletController(appCtx))
		walletEndpoint.Get("/", controllers.GetWalletBalanceController(appCtx))
		walletEndpoint.Patch("/", controllers.DisableWalletController(appCtx))
		walletEndpoint.Get("/transactions", controllers.ListWalletTransactionsController(appCtx))
		walletEndpoint.Post("/deposits", controllers.DepositWalletBalanceController(appCtx))
		walletEndpoint.Post("/withdrawals", controllers.WithdrawalWalletBalanceController(appCtx))
	}

	return
}