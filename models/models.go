package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserId    uuid.UUID `json:"userId" gorm:"primaryKey; not null;"`
	UserName  string    `json:"userName" gorm:"not null;"`
	UserEmail string    `json:"userEmail" gorm:"unique; not null" `
	Password  string    `json:"password" gorm:"not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Transaction struct {
	gorm.Model
	TransactionId uuid.UUID  `json:"transactionId"`
	Sender   uuid.UUID       `json:"senderId"`
	Receiver uuid.UUID       `json:"receiverId"`
	Amount   decimal.Decimal `json:"amount"`
	Message  string          `json:"message"`
}

type Account struct {
	gorm.Model
	AccountId    uuid.UUID       `json:"accountId" gorm:"primaryKey; unique; not null" `
	Balance      decimal.Decimal `json:"accountBalance" gorm:"not null;"`
	UserId       uuid.UUID       `json:"userId" gorm:"not null;"`
	Transactions []*Transaction  `json:"accountTransactions" gorm:"-"`
	Password     string
}
