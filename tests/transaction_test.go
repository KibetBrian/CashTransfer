package tests

import (
	"testing"

	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/services"
	"github.com/KibetBrian/fisa/utils"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T){
	accountId1, _ := uuid.FromString("d04253b5-e484-4874-886a-b1b471b64b6b")
	accountId2, _ := uuid.FromString("447e1d99-7992-4442-8330-ae9620e2f0ad")

	account1, success:= services.GetAccount(accountId1)
	require.True(t, success)

	account2, success := services.GetAccount(accountId2)
	require.True(t, success)

	routines := 10;
	isSuccessfulChan := make(chan bool)
	stringRes := make(chan string)
	transactionChan := make(chan *models.Transaction)
	amount := generate()

	for i:=0; i<routines; i++{
		go func(){
			transaction, res, isSuccessful :=services.DoubleEntry(accountId1, accountId2, decimal.NewFromInt32(amount))
  			isSuccessfulChan <- isSuccessful
			transactionChan <- transaction;
			stringRes <-res
		}()
	}
	
	for i:=0; i<routines; i++{
		successState := <- isSuccessfulChan
		transaction := <-transactionChan
		require.True(t, successState)
		require.NotEmpty(t,transaction)
		require.NotEmpty(t, transaction.Id)
		require.NotEmpty(t, transaction.Sender)
		require.NotEmpty(t, transaction.CreatedAt)
		require.NotEmpty(t, transaction.SenderAccountBalance)
		require.NotEmpty(t, transaction.ReceiverAccountBalance)
		require.Equal(t, transaction.Amount,decimal.NewFromInt32(amount))
		require.Equal(t, account1.Balance, transaction.SenderAccountBalance.Add(decimal.NewFromInt32(amount)))
		require.Equal(t, account2.Balance, transaction.ReceiverAccountBalance.Sub(decimal.NewFromInt32(amount)))
	}
}
//Generate random int32 between 1 and 1-
func generate () int32{
	num := utils.GenerateRandInt(1, 10)
	return int32(num)
}
