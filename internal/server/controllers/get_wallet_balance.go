package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetWalletBalanceController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		resp := models.SuccessResponse(map[string]interface{}{
			"balance": 0,
		})

		return ctx.
			Status(http.StatusOK).
			JSON(resp)
	}
}
