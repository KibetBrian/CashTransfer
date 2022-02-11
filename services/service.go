package services

import (
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

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

func DoubleEntry(senderAccountId uuid.UUID, receiverAccountId uuid.UUID, amount decimal.Decimal) (string, bool) {
	db, err := configs.ConnectDb()
	if err != nil {
		panic (err)
	}
	db.Begin();
	debitMessage, isDebitSuccessful := Debit(senderAccountId, amount)
	if !isDebitSuccessful {
		return debitMessage, false
	}
	_, isCreditSuccessful := Credit(receiverAccountId, amount)

	if !isCreditSuccessful{
		db.Rollback()
	}
	db.Commit()

	if isCreditSuccessful && isDebitSuccessful {
		return "Transaction Successful", true
	}
	return "Transaction Successful", true
}

func Deposit(accountId uuid.UUID, amount decimal.Decimal)(string, bool){
	var account models.Account 
	db, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}
	db.Where("account_id=?", accountId).First(&account);
	account.Balance=account.Balance.Add(amount);
	db.Save(&account)
	return "Deposit Successful", true
}
