package models

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/google/uuid"
)

type Transaction struct {
	ID           uuid.UUID              `json:"id"`
	WalletID     uuid.UUID              `json:"wallet_id"`
	Status       enum.TransactionStatus `json:"status"`
	TransactedAt time.Time              `json:"transacted_at"`
	Type         enum.TransactionType   `json:"type"`
	Amount       float64                `json:"amount"`
	ReferenceID  uuid.UUID              `json:"reference_id"`
}
