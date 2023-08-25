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
	{
		v1 := api.Group("/v1")
		{
			v1.Post("/init", controllers.InitAccountController(appCtx))

			v1Protected := v1.Use(middlewares.AuthorizationMiddleware(appCtx))
			{
				walletEp := v1Protected.Group("/wallet") // /wallet endpoint
				{
					walletEp.Post("/", controllers.EnableWalletController(appCtx))

					walletMustBeEnabledEp := walletEp.Use(middlewares.IsWalletEnabled(appCtx))
					{
						walletMustBeEnabledEp.Get("/", controllers.GetWalletController(appCtx))
						walletMustBeEnabledEp.Get("/transactions", controllers.ListWalletTransactionsController(appCtx))
						walletMustBeEnabledEp.Post("/deposits", controllers.DepositWalletBalanceController(appCtx))
						walletMustBeEnabledEp.Post("/withdrawals", controllers.WithdrawalWalletBalanceController(appCtx))
						walletMustBeEnabledEp.Patch("/", controllers.DisableWalletController(appCtx))
					}
				}
			}
		}
	}
	return
}
