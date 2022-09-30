package services

import (

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//Subtracts the amount from sender account
func Debit(senderAccountId uuid.UUID, amount decimal.Decimal, db *gorm.DB, tx *gorm.DB) (string, decimal.Decimal, bool) {
	var account models.Account

	tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("id=?", senderAccountId).First(&account)
	zero := decimal.NewFromInt(0)

	res := account.Balance.Sub(amount)
	if res.LessThan(zero) {
		tx.Rollback()
		return "Insufficient funds", account.Balance, false 
	}

	if amount.LessThanOrEqual(zero) {
		tx.Rollback()
		return "The amount to be sent is less than threshhold", account.Balance, false
	}

	account.Balance = account.Balance.Sub(amount)
	tx.Save(&account)

	return "Debit Successful", account.Balance, true
}

//Adds amount to the receivers account
func Credit(receiverAccountId uuid.UUID, amount decimal.Decimal, db *gorm.DB, tx *gorm.DB) (string, decimal.Decimal, bool) {

	var account models.Account
	tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("id=?", receiverAccountId).First(&account)

	account.Balance = amount.Add(account.Balance)
	tx.Save(&account)

	return "Credit successful", account.Balance, true
}

//Function to perform transaction
func DoubleEntry(senderAccountId uuid.UUID, receiverAccountId uuid.UUID, amount decimal.Decimal) (*models.Transaction, string, bool) {
	db, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}
	//Initialize transaction
	tx := db.Begin()
	defer func (){
		if r :=recover(); r != nil{
			tx.Rollback()
		}
	}()

	debitMessage, debitBalance, isDebitSuccessful := Debit(senderAccountId, amount, db, tx)
	if !isDebitSuccessful {
		return nil, debitMessage, false
	}
	_, creditBalance, isCreditSuccessful := Credit(receiverAccountId, amount, db, tx)

	if !isCreditSuccessful || !isDebitSuccessful {
		//Rollback if there was an error
		tx.Rollback()
		return nil, "Transaction failed", false
	}
	//Commmit transaction
	tx.Commit()
	transaction := &models.Transaction{
		Id:                     uuid.NewV4(),
		Amount:                 amount,
		Sender:                 senderAccountId,
		SenderAccountBalance:   debitBalance,
		Receiver:               receiverAccountId,
		ReceiverAccountBalance: creditBalance,
	}
	//Record the transactions to the database
	saveTransaction(transaction)

	if isCreditSuccessful && isDebitSuccessful {
		return transaction, "Transaction Successful", true
	}
	return transaction, "Transaction Successful", true
}

//Takes account id as input and adds amount to it
func Deposit(accountId uuid.UUID, amount decimal.Decimal) (string, bool) {
	var account models.Account

	db, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}

	db.Where("id=?", accountId).First(&account)
	account.Balance = account.Balance.Add(amount)
	db.Save(&account)

	transaction := &models.Transaction{
		Id:                     uuid.NewV4(),
		Receiver:               accountId,
		ReceiverAccountBalance: account.Balance,
		Amount:                 amount,
	}

	saveTransaction(transaction)

	return "Deposit Successful", true
}

//Saves transaction after an even of amount transanfer
func saveTransaction(t *models.Transaction) {
	db, err := configs.ConnectDb()
	if err != nil {
		panic("Failed to connect to the database")
	}
	
	db.Create(&t)
}
