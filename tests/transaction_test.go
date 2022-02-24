package tests

import (
	"fmt"
	"testing"

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

	for i:=0; i<routines; i++{
		go func(){
			res, isSuccessful :=services.DoubleEntry(accountId1, accountId2, decimal.NewFromInt32(generate()))
			isSuccessfulChan <- isSuccessful
			stringRes <-res
		}()
	}
	
	for i:=0; i<routines; i++{
		successState := <- isSuccessfulChan

		require.True(t, successState)
		
		fmt.Println(successState)
		res := <-stringRes
		fmt.Println(res)
	}
}
func generate () int32{
	num := utils.GenerateRandInt(1, 10)
	return int32(num)
}
