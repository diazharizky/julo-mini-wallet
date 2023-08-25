package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `form:"customer_xid" gorm:"column:id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
}
