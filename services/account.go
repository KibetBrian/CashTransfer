package services

import (
	"log"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

//Takes account id as input and return account object from the database
func GetAccount(id uuid.UUID) (*models.Account, bool) {
	var account models.Account

	db, err := configs.ConnectDb()
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	res := db.Where("id=?", id).First(&account)
	if res.RowsAffected < 1 {
		return nil, false
	}
	
	return &account, true
}

//Create account
func CreateAccount(userId uuid.UUID, tx *gorm.DB) (uuid.UUID, error) {
	var account models.Account
	
	account.Id = uuid.NewV4()
	account.Password, _ = utils.HashPassword(account.Password)
	account.HolderId = userId
	account.Balance=decimal.NewFromInt(0)

	res := tx.Create(&account)
	if res.Error != nil {
		return uuid.Nil, res.Error
	}

	return account.Id, nil
}
