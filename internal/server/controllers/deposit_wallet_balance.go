package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
)

func DepositWalletBalanceController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		resp := models.SuccessResponse("money has been added to the balance")

		return ctx.
			Status(http.StatusOK).
			JSON(resp)
	}
}
