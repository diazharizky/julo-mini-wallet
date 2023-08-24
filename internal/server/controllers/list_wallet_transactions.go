package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ListWalletTransactionsController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		accountID := ctx.Locals("account_id").(string)

		transactions, err := appCtx.ListWalletTransactionsModule.Call(
			uuid.MustParse(accountID),
		)
		if err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(models.FatalResponse())
		}

		resp := models.SuccessResponse(map[string]interface{}{
			"transactions": transactions,
		})

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
