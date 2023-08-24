package modules

import (
	"github.com/diazharizky/julo-mini-wallet/internal/app"
	"github.com/diazharizky/julo-mini-wallet/internal/models"
	"github.com/google/uuid"
)

type listWalletTransactionsModule struct {
	appCtx app.Ctx
}

func NewListWalletTransactionsModule(appCtx app.Ctx) listWalletTransactionsModule {
	return listWalletTransactionsModule{appCtx}
}

func (m listWalletTransactionsModule) Call(accountID uuid.UUID) ([]models.Transaction, error) {
	wallet, err := m.appCtx.WalletRepository.GetByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	transactions, err := m.appCtx.TransactionRepository.List(wallet.ID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
