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
		require.NotEmpty(t, transaction.Sender)
		require.NotEmpty(t, transaction.CreatedAt)
		require.NotEmpty(t, transaction.Id)
		require.Equal(t, transaction.Amount,decimal.NewFromInt32(amount))
		require.Equal(t, transaction.Sender, transaction.Receiver)	
	}
}
func generate () int32{
	num := utils.GenerateRandInt(1, 10)
	return int32(num)
}
