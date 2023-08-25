package middlewares

import (
	"net/http"
	"strings"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthorizationMiddleware(appCtx app.Ctx) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		authHeader, ok := ctx.GetReqHeaders()["Authorization"]
		if !ok {
			return ctx.SendStatus(http.StatusUnauthorized)
		}

		val := strings.Split(authHeader, " ")
		if len(val) != 2 {
			return ctx.SendStatus(http.StatusUnauthorized)
		}

		accountID, err := utils.Decode(val[1])
		if err != nil {
			return ctx.SendStatus(http.StatusUnauthorized)
		}

		ctx.Locals("account_id", *accountID)

		return ctx.Next()
	}
}
