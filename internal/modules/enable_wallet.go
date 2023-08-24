package modules

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type enableWalletModule struct {
	appCtx app.Ctx
}

func NewEnableWalletModule(appCtx app.Ctx) enableWalletModule {
	return enableWalletModule{appCtx}
}

func (m enableWalletModule) Call(accountID uuid.UUID) (*models.Wallet, error) {
	wallet, err := m.appCtx.WalletRepository.GetByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	if wallet.Status == enum.WalletStatusEnabled {
		return nil, app.WalletIsAlreadyEnabled
	}

	if err := m.appCtx.WalletRepository.Enable(wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}
