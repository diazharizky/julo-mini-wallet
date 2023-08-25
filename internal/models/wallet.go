package models

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/google/uuid"
)

type Wallet struct {
	ID         uuid.UUID         `json:"id" gorm:"column:id;primaryKey"`
	OwnedBy    uuid.UUID         `json:"owned_by" gorm:"column:owned_by;not null;foreignKey"`
	Status     enum.WalletStatus `json:"status" gorm:"column:status;not null"`
	EnabledAt  *time.Time        `json:"enabled_at,omitempty" gorm:"column:enabled_at"`
	DisabledAt *time.Time        `json:"disabled_at,omitempty" gorm:"column:disabled_at"`
	Balance    float64           `json:"balance" gorm:"column:balance;not null"`
}

func NewDefaultWAllet(accountID uuid.UUID) *Wallet {
	return &Wallet{
		ID:      uuid.New(),
		OwnedBy: accountID,
		Status:  enum.WalletStatusDisabled,
		Balance: 0,
	}
}
