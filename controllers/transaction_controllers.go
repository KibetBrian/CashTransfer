package controllers

import (
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/services"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

//Finds account id associated with particular email
func findAccountId(email string) (uuid.UUID, bool) {
	var user models.User
	db, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}
	res := db.Where("email = ?", email).First(&user)
	if res.RowsAffected < 1 {
		return uuid.Nil, false
	}
	return user.AccountId, true
}

//Adds amount to the account
func Deposit(c *gin.Context) {
	var TransactionReq models.TransactionRequest
	var Transaction models.Transaction
	id := uuid.NewV4()
	c.ShouldBindJSON(&TransactionReq)
	Transaction.Id = id
	receiverAccountId, isValid := findAccountId(TransactionReq.ReceiverEmail)
	if !isValid {
		c.JSON(200, "Seems we don't have users with that email")
		return
	}
	Transaction.Receiver = receiverAccountId
	//Call services deposit func with credential and amount
	message, isSuccessful := services.Deposit(receiverAccountId, TransactionReq.Amount)
	if isSuccessful {
		c.JSON(200, gin.H{"Message": message, "Transaction": TransactionReq})
		return
	}

}

//Transfer amount to another account
func Send (c *gin.Context){
	var TransactionReq models.TransactionRequest
	c.ShouldBindJSON(&TransactionReq);

	receiverAccountId, isValid := findAccountId(TransactionReq.ReceiverEmail)
	if !isValid{
		c.JSON(404, gin.H{"Message":"It seems we don't have a user with that email", "Email":TransactionReq.ReceiverEmail})
		return
	}

	senderAccountId, isValid := findAccountId(TransactionReq.SenderEmail)
	if !isValid {
		c.JSON(403, gin.H{"Message":"Check credential and try again"})
		return
	}

	if receiverAccountId == senderAccountId{
		c.JSON(403, gin.H{"Message": "You cannot send money to yourself"})
		return
	}
	//If the credentials are correct	, send the account ids' to services package for processing
	transaction, message, successful:=services.DoubleEntry(senderAccountId, receiverAccountId,TransactionReq.Amount);
	if !successful{
		c.JSON(403, message);
		return
	}	
	c.JSON(200, gin.H{"Message": message, "Transaction: ": transaction})
}
