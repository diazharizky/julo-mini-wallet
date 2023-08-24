package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
)

func InitAccountController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var newAccount models.Account

		if err := ctx.BodyParser(&newAccount); err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(models.HTTPResponse{
				Status: "failed",
				Data: map[string]interface{}{
					"message": "Failed to parse body",
				},
			})
		}

		token := appCtx.GenerateTokenModule.Call(newAccount.ID)

		resp := models.SuccessResponse(map[string]interface{}{
			"token": token,
		})

		return ctx.
			Status(http.StatusOK).
			JSON(resp)
	}
}
