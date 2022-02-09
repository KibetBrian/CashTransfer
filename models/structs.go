package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)


type Transaction struct{
	gorm.Model
	FROM uuid.UUID
	TO uuid.UUID
	AMOUNT decimal.Decimal
	Message string
}

type Account struct {
	gorm.Model
	AccountId  uuid.UUID
	Balance  decimal.Decimal
	Holder uuid.UUID
	Transactions []Transaction 
}