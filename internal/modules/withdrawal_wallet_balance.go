package modules

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type withdrawalWalletBalanceModule struct {
	appCtx app.Ctx
}

func NewWithdrawalWalletBalanceModule(appCtx app.Ctx) withdrawalWalletBalanceModule {
	return withdrawalWalletBalanceModule{appCtx}
}

func (m withdrawalWalletBalanceModule) Call(newWithdrawal *models.Withdrawal) error {
	newWithdrawal.ID = uuid.New()
	newWithdrawal.WithdrawnAt = time.Now()
	newWithdrawal.Status = enum.WithdrawalStatusSuccess

	return nil
}
