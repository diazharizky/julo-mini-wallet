package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DepositWalletBalanceController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var newDeposit models.Deposit

		if err := ctx.BodyParser(&newDeposit); err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(models.FailedResponse(map[string]interface{}{
					"message": "Failed to parse body",
				}))
		}

		accountID := ctx.Locals("account_id").(string)
		newDeposit.DepositedBy = uuid.MustParse(accountID)

		if err := appCtx.DepositWalletBalanceModule.Call(&newDeposit); err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(models.FatalResponse())
		}

		resp := models.SuccessResponse(map[string]interface{}{
			"deposit": newDeposit,
		})

		return ctx.
			Status(http.StatusOK).
			JSON(resp)
	}
}
