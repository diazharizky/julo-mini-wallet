package modules

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/google/uuid"
)

type initializeAccountModule struct {
	appCtx app.Ctx
}

func NewInitializeAccountModule(appCtx app.Ctx) initializeAccountModule {
	return initializeAccountModule{appCtx}
}

func (m initializeAccountModule) Call(customerXID uuid.UUID) {

}
