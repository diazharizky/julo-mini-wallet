package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type disableWalletParams struct {
	IsDisable bool `form:"is_disable"`
}

func DisableWalletController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var params disableWalletParams

		if err := ctx.BodyParser(&params); err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(models.FailedResponse(map[string]interface{}{
					"message": "Failed to parse body",
				}))
		}

		accountID := ctx.Locals("account_id").(string)
		wallet, err := appCtx.DisableWalletModule.Call(
			uuid.MustParse(accountID),
		)
		if err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(models.FatalResponse())
		}

		resp := models.SuccessResponse(map[string]interface{}{
			"wallet": wallet,
		})

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
