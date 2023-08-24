package models

import "github.com/google/uuid"

type Account struct {
	ID uuid.UUID `form:"customer_xid" gorm:"primaryKey;column:id"`
}
