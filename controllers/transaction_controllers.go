package controllers

import (
	"fmt"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
//Finds account id associated with particular email
func findAccountId(email string, Transaction *models.Transaction) uuid.UUID {
	var user models.User
	db, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}
	res := db.Where("user_email = ?", Transaction.Receiver).First(&user)
	if res.RowsAffected < 1 {
		fmt.Println("No  user found")
	}
	return user.AccountId;
}

func Deposit(c *gin.Context) {
	var TransactionReq models.TransactionRequest
	var Transaction models.Transaction
	id := uuid.New()
	c.ShouldBindJSON(TransactionReq)
	Transaction.TransactionId = id
	receiverAccountId:=findAccountId(TransactionReq.ReceiverEmail, &Transaction)
	Transaction.Receiver=receiverAccountId
	//Call services deposit func with credetial and amoun
	message, isSuccessful :=services.Deposit(receiverAccountId, Transaction.Amount)
	if isSuccessful{
		c.JSON(200, gin.H{"Message": message, "Transaction":Transaction })
	}
	
}
