package tests

import (
	"log"
	"testing"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

//Test user creation with random data
func TestUserRegistration(t *testing.T){
	testUser := &models.User{
		Name: utils.GenerateRandomUserName(),
		Email: utils.GenerateRandomEmail(),
		Password: utils.GenerateRandomPassword(6,48),
		Id: uuid.New(),
	}
	db, err := configs.ConnectDb()
	
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	res :=db.Create(&testUser)

	require.NoError(t, res.Error)	
	require.NotEmpty(t, res.RowsAffected)

	require.NotEmpty(t, testUser.Id)
	require.NotEmpty(t, testUser.Email)
	require.NotEmpty(t, testUser.Password)
	require.NotEmpty(t, testUser.Name)

}