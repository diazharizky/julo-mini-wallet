package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/diazharizky/julo-mini-wallet/pkg/utils"
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
				JSON(models.FailedParseBody())
		}

		if err := appCtx.InitAccountModule.Call(params.ID); err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(models.FatalResponse())
		}

		token := utils.EncodeToString(params.ID.String())
		resp := models.SuccessResponse(map[string]interface{}{
			"token": token,
		})

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
