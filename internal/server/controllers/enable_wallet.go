package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
)

func EnableWalletController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		accountID := ctx.Locals("account_id")

		resp := models.SuccessResponse(map[string]interface{}{
			"message":   "Wallet enabled",
			"accountId": accountID,
		})

		return ctx.
			Status(http.StatusOK).
			JSON(resp)
	}
}
