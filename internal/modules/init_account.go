package modules

import (
	"fmt"

	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type initAccountModule struct {
	appCtx app.Ctx
}

func NewInitAccountModule(appCtx app.Ctx) initAccountModule {
	return initAccountModule{appCtx}
}

func (m initAccountModule) Call(id uuid.UUID) error {
	tx := m.appCtx.AccountRepository.BeginTx()

	newAccount := &models.Account{ID: id}
	if err := m.appCtx.AccountRepository.Create(tx, newAccount); err != nil {
		return err
	}

	_, err := m.appCtx.WalletRepository.Create(tx, newAccount.ID)
	if err != nil {
		if err := m.appCtx.WalletRepository.RollbackTx(tx); err != nil {
			fmt.Printf("rollback transaction error: %v\n", err)
		}

		return err
	}

	if err := m.appCtx.WalletRepository.CommitTx(tx); err != nil {
		fmt.Printf("commit transaction error: %v\n", err)
	}

	return nil
}
