package controllers

import (
	"net/http"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/diazharizky/julo-mini-wallet/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func InitAccountController(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var newAccount models.Account
		if err := ctx.BodyParser(&newAccount); err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(models.FailedParseBody())
		}

		if err := appCtx.InitAccountModule.Call(newAccount.ID); err != nil {
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(models.FatalResponse())
		}

		token := utils.Encode(newAccount.ID.String())
		resp := models.SuccessResponse(map[string]interface{}{
			"token": token,
		})

		return ctx.Status(http.StatusOK).JSON(resp)
	}
}
