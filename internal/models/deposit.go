package models

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/google/uuid"
)

type Deposit struct {
	ID          uuid.UUID          `json:"id"`
	DepositedBy uuid.UUID          `json:"deposited_by"`
	Status      enum.DepositStatus `json:"status"`
	DepositedAt time.Time          `json:"deposited_at"`
	Amount      float64            `form:"amount" json:"amount"`
	ReferenceID uuid.UUID          `form:"reference_id" json:"reference_id"`
}
