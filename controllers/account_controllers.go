package controllers

import (
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAccount(c *gin.Context) {
	var Account models.Account
	c.ShouldBindJSON(&Account)
	id := uuid.New()
	Account.AccountId = id
	Account.Password, _ =utils.HashPassword(Account.Password)
	db, err := configs.ConnectDb()
	if err != nil {
		c.JSON(500, "Error connecting to the database")
		return
	}
	db.Create(&Account)
}

func DeleteAccount(c *gin.Context) {

}
