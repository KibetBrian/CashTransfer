package services

import (
	"log"
	"testing"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

//Test user creation with random data
func TestRegisterUser(t *testing.T){
	testUser := &models.User{
		Name: utils.GenerateRandomUserName(),
		Email: utils.GenerateRandomEmail(),
		Password: utils.GenerateRandomPassword(6,48),
		Id: uuid.NewV4(),
	}

	db, err := configs.ConnectDb()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	isValid := utils.CheckValidity(testUser.Email)
	res :=db.Create(&testUser)
	require.True(t, isValid)
	require.NoError(t, res.Error)	
	require.NotEmpty(t, res.RowsAffected)
	require.NotEmpty(t, testUser.Id)
	require.NotEmpty(t, testUser.Email)
	require.NotEmpty(t, testUser.Password)
	require.NotEmpty(t, testUser.Name)
}
