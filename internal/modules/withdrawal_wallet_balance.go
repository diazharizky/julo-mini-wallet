package modules

import (
	"fmt"

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

	tx := m.appCtx.WalletRepository.BeginTx()
	if err := m.appCtx.WalletRepository.DepositBalance(tx, wallet, -withdrawalAmount); err != nil {
		return err
	}

	newTrx := models.Transaction{
		WalletID:    wallet.ID,
		Type:        enum.TransactionTypeWithdrawal,
		Amount:      withdrawalAmount,
		ReferenceID: newWithdrawal.ReferenceID,
	}
	if err := m.appCtx.TransactionRepository.Create(tx, &newTrx); err != nil {
		if err := m.appCtx.TransactionRepository.RollbackTx(tx); err != nil {
			fmt.Printf("rollback transaction error: %v\n", err)
		}

		return err
	}

	if err := m.appCtx.TransactionRepository.CommitTx(tx); err != nil {
		fmt.Printf("commit transaction error: %v\n", err)
	}

	newWithdrawal.ID = newTrx.ID
	newWithdrawal.Status = enum.WithdrawalStatusSuccess
	newWithdrawal.WithdrawnAt = newTrx.TransactedAt

	return nil
}
