package services

import (
	"fmt"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

//Subtracts the amount from sender accounts
func Debit(senderAccountId uuid.UUID, amount decimal.Decimal) (string, decimal.Decimal, bool) {
	db, err := configs.ConnectDb()
	var account models.Account
	if err != nil {
		panic(err)
	}
	db.Where("id=?", senderAccountId).First(&account)
	fmt.Println(senderAccountId)
	zero := decimal.NewFromInt(0)
	res := account.Balance.Sub(amount)
	if res.LessThan(zero) {
		return "Insufficient funds", account.Balance, false
	}
	if account.Balance.LessThan(zero) {
		return "You can only transfer amount grater than zero", account.Balance, false
	}
	account.Balance = account.Balance.Sub(amount)
	db.Save(&account)
	return "Debit Successful", account.Balance, true
}

//Adds amount to the receivers account
func Credit(receiverAccountId uuid.UUID, amount decimal.Decimal) (string, decimal.Decimal, bool) {
	var account models.Account

	db, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}
	db.Where("account_id=?", receiverAccountId).First(&account)
	account.Balance = amount.Add(account.Balance)
	db.Save(&account)
	return "Credit successful", account.Balance, true
}

//Function to perform transaction
func DoubleEntry(senderAccountId uuid.UUID, receiverAccountId uuid.UUID, amount decimal.Decimal) (string, bool) {
	db, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}
	//Initialize transactions
	db.Begin()
	debitMessage, debitBalance, isDebitSuccessful := Debit(senderAccountId, amount)
	if !isDebitSuccessful {
		return debitMessage, false
	}
	_, creditBalance, isCreditSuccessful := Credit(receiverAccountId, amount)

	if !isCreditSuccessful {
		//Rollback if there was an error
		db.Rollback()
	}
	//Commmit transaction
	db.Commit()
	transaction := &models.Transaction{
		Id:                     uuid.New(),
		Amount:                 amount,
		Sender:                 senderAccountId,
		SenderAccountBalance:   debitBalance,
		Receiver:               receiverAccountId,
		ReceiverAccountBalance: creditBalance,
	}
	//Record the transactions to the database
	saveTransaction(transaction)
	if isCreditSuccessful && isDebitSuccessful {
		return "Transaction Successful", true
	}
	return "Transaction Successful", true
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
		Id:                     uuid.New(),
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
	db.Save(&t)
}
