package handlers

import (
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	uuid "github.com/satori/go.uuid"
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
