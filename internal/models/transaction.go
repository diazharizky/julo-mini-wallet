package models

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/google/uuid"
)

type Transaction struct {
	ID           uuid.UUID              `json:"id" gorm:"column:id;primaryKey"`
	WalletID     uuid.UUID              `json:"wallet_id" gorm:"column:wallet_id;not null;foreignKey"`
	Status       enum.TransactionStatus `json:"status" gorm:"column:status;not null"`
	TransactedAt time.Time              `json:"transacted_at" gorm:"column:transacted_at;not null"`
	Type         enum.TransactionType   `json:"type" gorm:"column:type;not null"`
	Amount       float64                `json:"amount" gorm:"column:amount;not null"`
	ReferenceID  uuid.UUID              `json:"reference_id" gorm:"column:reference_id;not null;unique"`
}
