package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
)

func ListWalletTransactionsController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		resp := models.SuccessResponse("list of transactions")

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
