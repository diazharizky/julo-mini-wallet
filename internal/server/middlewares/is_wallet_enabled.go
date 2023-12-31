package middlewares

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func IsWalletEnabled(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		accountID := ctx.Locals("account_id").(string)
		wallet, err := appCtx.WalletRepository.GetByAccountID(
			uuid.MustParse(accountID),
		)
		if err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(models.FatalResponse())
		}

		if wallet == nil {
			return ctx.
				Status(http.StatusNotFound).
				JSON(models.FailedResponse(map[string]interface{}{
					"error": "app: data cannot be found",
				}))
		}

		if wallet.Status == enum.WalletStatusDisabled {
			return ctx.
				Status(http.StatusConflict).
				JSON(models.FailedResponse(map[string]interface{}{
					"error": "app: wallet is disabled",
				}))
		}

		return ctx.Next()
	}
}
