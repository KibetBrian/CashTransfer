package models

import (
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type AccountReq struct {
	AccountId uuid.UUID `json:"accountId" binding:"required"`
}

type User struct {
	gorm.Model
	Id        uuid.UUID `json:"id" gorm:"primaryKey; not null;"`
	Name      string    `json:"name" gorm:"not null;" binding:"required"`
	Email     string    `json:"email" gorm:"unique; not null" binding:"required" `
	Password  string    `json:"password" gorm:"not null;" binding:"required"`
	CreatedAt time.Time `json:"createdAt" sql:"type:timestamp" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
	AccountId uuid.UUID `json:"accountId" gorm:"unique; not null"`
}

type Transaction struct {
	gorm.Model
	Id                     uuid.UUID       `json:"transactionId"`
	Sender                 uuid.UUID       `json:"senderId"`
	SenderAccountBalance   decimal.Decimal `json:"senderAccountBalance"`
	Receiver               uuid.UUID       `json:"receiverId"`
	ReceiverAccountBalance decimal.Decimal `json:"receiverAccountBalance"`
	Amount                 decimal.Decimal `json:"amount"`
	Message                string          `json:"message"`
}

type Account struct {
	gorm.Model
	Id           uuid.UUID       `json:"accountId" gorm:"primaryKey; unique; not null" `
	Balance      decimal.Decimal `json:"accountBalance" gorm:"not null;"`
	HolderId     uuid.UUID       `json:"holderId" gorm:"not null;"`
	Transactions []*Transaction  `json:"accountTransactions" gorm:"-"`
	Password     string          `json:"password" gorm:"not null;"`
}

type TransactionRequest struct {
	SenderEmail    string          `json:"senderEmail"`
	SenderPassword string          `json:"senderPassword"`
	ReceiverEmail  string          `json:"receiverEmail"`
	Amount         decimal.Decimal `json:"amount" binding:"required,gt=0"`
	Message        string          `json:"message"`
}

type LoginRequest struct{
	Email string 	`json:"email"`
	Password string  `json:"password"`
}

type LoginResponse struct{
	AccessToken string `json:"accessToken"`
	AccessTokenExpiresAt time.Time `json:"accessTokenExpiresAt"`
	RefreshToken string `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time `json:"refreshtTokenExpiresAt"`
	User  User `json:"userData"`
}

type RefreshTokenReq struct{
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenStored struct{
	TokenId uuid.UUID `json:"tokenId"`
	RefreshToken string `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time `json:"refreshtTokenExpiresAt"`
	ClientIp string `json:"clientIp"`
	UserAgent string `json:"userAgent"`
}

type RefreshTokenResponse struct{
	AccessToken string `json:"accessToken"`
	AccessTokenExpiresAt time.Time `json:"accessTokenExpiresAt"`
	RefreshToken string `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time `json:"refreshtTokenExpiresAt"`
}
