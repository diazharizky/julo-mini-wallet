package modules

import "github.com/diazharizky/julo-mini-wallet/internal/app"

type enableWalletModule struct {
	appCtx app.Ctx
}

func NewEnableWalletModule(appCtx app.Ctx) enableWalletModule {
	return enableWalletModule{appCtx}
}

func (m enableWalletModule) Call() {}
