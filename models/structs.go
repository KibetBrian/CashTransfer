package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
	"github.com/ubgo/gormuuid"
)


type User struct {
	gorm.Model
	UserId       uuid.UUID          `json:"userId" gorm:"primaryKey; not null;"`
	UserName  string             `json:"userName"`
	UserEmail string             `gorm:"unique" json:"userEmail"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Accounts  gormuuid.UUIDArray `gorm:"type:uuid[]"`
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
	User User `gorm:"foreignKey: UserId"`
	AccountId  uuid.UUID `json:"accountId" gorm:"primaryKey; unique; not null" `
	Balance  decimal.Decimal `json:"accountBalance"`
	UserId uuid.UUID 	`json:"userId" gorm:"not null;"`
	Transactions [] *Transaction `json:"accountTransactions" gorm:"-"`
	Password  string
}