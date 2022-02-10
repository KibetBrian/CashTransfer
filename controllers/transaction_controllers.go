package controllers

import (
	"fmt"

	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/configs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func findAccountId(email string,Transaction *models.Transaction) {
	var user models.User
	db, err := configs.ConnectDb();
	if err != nil {
		panic(err)
	}
	res := db.Where("user_email = ?",Transaction.Receiver).First(&user)
	if res.RowsAffected <1{
		fmt.Println("No  user found")
	}
}

func Deposit(c *gin.Context){
	var Transaction models.Transaction;
	id := uuid.New();
	c.ShouldBindJSON(Transaction);
	Transaction.TransactionId =id;
	
	fmt.Println(Transaction);
}