package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type initAccountParams struct {
	ID uuid.UUID `form:"customer_xid"`
}

func InitAccountController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var params initAccountParams

		if err := ctx.BodyParser(&params); err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(models.FailedResponse(map[string]interface{}{
					"message": "Failed to parse body",
				}))
		}

		token := appCtx.GenerateTokenModule.Call(params.ID)

		resp := models.SuccessResponse(map[string]interface{}{
			"token": token,
		})

		return ctx.
			Status(http.StatusOK).
			JSON(resp)
	}
}
