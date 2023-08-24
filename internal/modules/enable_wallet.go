package modules

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
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
	existingWallet, err := m.appCtx.WalletRepository.GetByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	if existingWallet != nil {
		return nil, app.WalletIsAlreadyEnabled
	}

	newWallet, err := m.appCtx.WalletRepository.Create(accountID)
	if err != nil {
		return nil, err
	}

	return newWallet, nil
}
