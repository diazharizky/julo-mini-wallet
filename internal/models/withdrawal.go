package models

import (
	"time"

	"github.com/diazharizky/julo-mini-wallet/internal/enum"
	"github.com/google/uuid"
)

type Withdrawal struct {
	ID          uuid.UUID             `json:"id"`
	WithdrawnBy uuid.UUID             `json:"withdrawn_by"`
	Status      enum.WithdrawalStatus `json:"status"`
	WithdrawnAt time.Time             `json:"withdrawn_at"`
	Amount      float64               `form:"amount" json:"amount"`
	ReferenceID uuid.UUID             `form:"reference_id" json:"reference_id"`
}
