package modules

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type depositWalletBalanceModule struct {
	appCtx app.Ctx
}

func NewDepositWalletBalanceModule(appCtx app.Ctx) depositWalletBalanceModule {
	return depositWalletBalanceModule{appCtx}
}

func (m depositWalletBalanceModule) Call(newDeposit *models.Deposit) error {
	newDeposit.ID = uuid.New()
	newDeposit.DepositedAt = time.Now()
	newDeposit.Status = enum.DepositStatusSuccess

	return nil
}
