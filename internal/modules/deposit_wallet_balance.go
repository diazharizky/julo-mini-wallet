package modules

import (
	"fmt"

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

	tx := m.appCtx.WalletRepository.BeginTx()

	depositAmount := newDeposit.Amount
	if err := m.appCtx.WalletRepository.DepositBalance(tx, wallet, depositAmount); err != nil {
		return err
	}

	newTrx := models.Transaction{
		WalletID:    wallet.ID,
		Type:        enum.TransactionTypeDeposit,
		Amount:      depositAmount,
		ReferenceID: newDeposit.ReferenceID,
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

	newDeposit.ID = newTrx.ID
	newDeposit.Status = enum.DepositStatusSuccess
	newDeposit.DepositedAt = newTrx.TransactedAt

	return nil
}
