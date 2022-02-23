package controllers

import (
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func CreateAccount(c *gin.Context) {
	var Account models.Account
	var user models.User
	c.ShouldBindJSON(&Account)
	id := uuid.NewV4()
	Account.Id = id
	Account.Password, _ = utils.HashPassword(Account.Password)
	db, err := configs.ConnectDb()
	if err != nil {
		c.JSON(500, "Error connecting to the database")
		return
	}

	res := db.Where("id=?",Account.HolderId ).First(&user)
	if res.RowsAffected < 1 {
		c.JSON(404, gin.H{"Message": "No such user", "User":user})
		return
	}

	db.AutoMigrate(&models.Account{})
	db.Create(&Account)
	db.Model(&user).Update("account_id", Account.Id)
	c.JSON(200, gin.H{"Message": "Account Created", "Account": Account})
}

func DeleteAccount(c *gin.Context) {

}
