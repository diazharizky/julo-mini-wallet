package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func EnableWalletController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		accountID := ctx.Locals("account_id").(string)
		wallet, err := appCtx.EnableWalletModule.Call(
			uuid.MustParse(accountID),
		)
		if err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(models.FailedResponse(map[string]interface{}{
					"error": err.Error(),
				}))
		}

		resp := models.SuccessResponse(map[string]interface{}{
			"wallet": wallet,
		})

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
