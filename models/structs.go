package models

import (
	"time"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id int
	UserId     uuid.UUID 
	Username   string `json:"userName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	CreatedAt time.Time 
	UpdatedAt time.Time
	DeletedAt time.Time
	Accounts  []uuid.UUID `gorm:"type:array"`
}

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