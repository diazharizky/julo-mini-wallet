package modules

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
)

type depositWalletBalanceModule struct {
	appCtx app.Ctx
}

func NewDepositWalletBalanceModule(appCtx app.Ctx) depositWalletBalanceModule {
	return depositWalletBalanceModule{appCtx}
}

func (m depositWalletBalanceModule) Call(newDeposit *models.Deposit) error {
	wallet, err := m.appCtx.WalletRepository.GetByAccountID(newDeposit.DepositedBy)
	if err != nil {
		return err
	}

	depositAmount := newDeposit.Amount
	if err := m.appCtx.WalletRepository.Deposit(wallet, depositAmount); err != nil {
		return err
	}

	newTrx := models.Transaction{
		WalletID:    wallet.ID,
		Type:        enum.TransactionTypeDeposit,
		Amount:      depositAmount,
		ReferenceID: newDeposit.ReferenceID,
	}

	if err := m.appCtx.TransactionRepository.Create(&newTrx); err != nil {
		return err
	}

	newDeposit.ID = newTrx.ID
	newDeposit.Status = enum.DepositStatusSuccess
	newDeposit.DepositedAt = newTrx.TransactedAt

	return nil
}
