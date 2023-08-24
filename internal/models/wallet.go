package models

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/google/uuid"
)

type Wallet struct {
	ID        uuid.UUID         `json:"id" gorm:"primaryKey;column:id"`
	OwnedBy   uuid.UUID         `json:"owned_by" gorm:"column:owned_by;not null"`
	Status    enum.WalletStatus `json:"status" gorm:"column:status;not null"`
	EnabledAt time.Time         `json:"enabled_at" gorm:"column:enabled_at;not null"`
	Balance   float64           `json:"balance" gorm:"column:balance;not null"`
}

func NewDefaultWAllet(accountID uuid.UUID) *Wallet {
	return &Wallet{
		ID:        uuid.New(),
		OwnedBy:   accountID,
		Status:    enum.WalletStatusEnabled,
		EnabledAt: time.Now(),
		Balance:   0,
	}
}
