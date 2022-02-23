package services

import (
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

//Subtracts the amount from sender accounts
func Debit(senderAccountId uuid.UUID, amount decimal.Decimal) (string, bool) {
	db, err := configs.ConnectDb()
	var account models.Account
	if err != nil {
		panic(err)
	}
	db.Where("account_id=?", senderAccountId).First(&account)
	zero := decimal.NewFromInt(0)
	res := account.Balance.Sub(amount)
	if res.LessThan(zero) {
		return "Insufficient funds", false
	}
	account.Balance = account.Balance.Sub(amount)
	db.Save(&account)
	return "Successful Debit", true
}

//Adds amount to the receivers account
func Credit(receiverAccountId uuid.UUID, amount decimal.Decimal) (string, bool) {
	var account models.Account

	db, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}

	db.Where("account_id=?", receiverAccountId).First(&account)
	account.Balance = amount.Add(account.Balance)
	db.Save(&account)

	return "Credit Successful", true
}

//Function to perform transaction
func DoubleEntry(senderAccountId uuid.UUID, receiverAccountId uuid.UUID, amount decimal.Decimal) (string, bool) {
	db, err := configs.ConnectDb()
	if err != nil {
		panic (err)
	}
	//Initialize transactions
	db.Begin();
	debitMessage, isDebitSuccessful := Debit(senderAccountId, amount)
	if !isDebitSuccessful {
		return debitMessage, false
	}
	_, isCreditSuccessful := Credit(receiverAccountId, amount)

	if !isCreditSuccessful{
		//Rollback if there was an error
		db.Rollback()
	}
	//Commmit transaction
	db.Commit()

	if isCreditSuccessful && isDebitSuccessful {
		return "Transaction Successful", true
	}
	return "Transaction Successful", true
}

//Adds amount to the account id received
func Deposit(accountId uuid.UUID, amount decimal.Decimal)(string, bool){
	var account models.Account 
	db, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}
	db.Where("id=?", accountId).First(&account);
	account.Balance=account.Balance.Add(amount);
	db.Save(&account)
	return "Deposit Successful", true
}
