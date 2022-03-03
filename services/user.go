package services

import (
	"log"
	"time"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	uuid "github.com/satori/go.uuid"
)

func RegisterUser(user *models.User) (*models.User, error) {
	db, err := configs.ConnectDb()
	if err != nil {
		log.Fatal("Error connecting database...Err: ", err)
		return nil, err
	}
	user.Id = uuid.NewV4()
	user.CreatedAt = time.Now()
	tx := db.Begin()
	accountId, err := CreateAccount(user.Id, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	user.AccountId=accountId
	result := tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return  user, nil
}
