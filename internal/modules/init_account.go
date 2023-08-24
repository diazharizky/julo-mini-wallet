package modules

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/google/uuid"
)

type initAccountModule struct {
	appCtx app.Ctx
}

func NewInitAccountModule(appCtx app.Ctx) initAccountModule {
	return initAccountModule{appCtx}
}

func (m initAccountModule) Call(id uuid.UUID) error {
	_, err := m.appCtx.WalletRepository.Create(id)
	if err != nil {
		return err
	}

	return nil
}
