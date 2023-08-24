package modules

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type disableWalletModule struct {
	appCtx app.Ctx
}

func NewDisableWalletModule(appCtx app.Ctx) disableWalletModule {
	return disableWalletModule{appCtx}
}

func (m disableWalletModule) Call(accountID uuid.UUID) (*models.Wallet, error) {
	wallet, err := m.appCtx.WalletRepository.GetByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	if wallet == nil {
		return nil, app.WalletIsDisabled
	}

	if err := m.appCtx.WalletRepository.Disable(wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}
