package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
)

func DisableWalletController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		resp := models.SuccessResponse("successfully disable wallet")

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
