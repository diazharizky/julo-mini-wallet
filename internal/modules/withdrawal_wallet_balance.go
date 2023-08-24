package modules

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
)

type withdrawalWalletBalanceModule struct {
	appCtx app.Ctx
}

func NewWithdrawalWalletBalanceModule(appCtx app.Ctx) withdrawalWalletBalanceModule {
	return withdrawalWalletBalanceModule{appCtx}
}

func (m withdrawalWalletBalanceModule) Call(newWithdrawal *models.Withdrawal) error {
	wallet, err := m.appCtx.WalletRepository.GetByAccountID(newWithdrawal.WithdrawnBy)
	if err != nil {
		return err
	}

	withdrawalAmount := newWithdrawal.Amount
	if wallet.Balance < withdrawalAmount {
		return app.InsufficientBalance
	}

	if err := m.appCtx.WalletRepository.Deposit(wallet, -withdrawalAmount); err != nil {
		return err
	}

	newTrx := models.Transaction{
		WalletID:    wallet.ID,
		Type:        enum.TransactionTypeWithdrawal,
		Amount:      withdrawalAmount,
		ReferenceID: newWithdrawal.ReferenceID,
	}
	if err := m.appCtx.TransactionRepository.Create(&newTrx); err != nil {
		return err
	}

	newWithdrawal.ID = newTrx.ID
	newWithdrawal.Status = enum.WithdrawalStatusSuccess
	newWithdrawal.WithdrawnAt = newTrx.TransactedAt

	return nil
}
