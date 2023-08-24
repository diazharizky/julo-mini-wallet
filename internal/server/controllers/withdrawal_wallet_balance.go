package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func WithdrawalWalletBalanceController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var newWithdrawal models.Withdrawal

		if err := ctx.BodyParser(&newWithdrawal); err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(models.FailedParseBody())
		}

		accountID := ctx.Locals("account_id").(string)
		newWithdrawal.WithdrawnBy = uuid.MustParse(accountID)

		if err := appCtx.WithdrawalWalletBalanceModule.Call(&newWithdrawal); err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(models.FailedResponse(map[string]interface{}{
					"error": err.Error(),
				}))
		}

		resp := models.SuccessResponse(map[string]interface{}{
			"withdrawal": newWithdrawal,
		})

		return ctx.
			Status(http.StatusOK).
			JSON(resp)
	}
}
