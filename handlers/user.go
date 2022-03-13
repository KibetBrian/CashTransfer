package handlers

import (
	"net/http"
	"time"

	"github.com/KibetBrian/fisa/auth"
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/database/redis"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/services"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Hello"})
}

//Gets user input, validates and adds to the database
func RegisterUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrResponse(err))
		return
	}
	isValid := utils.CheckValidity(user.Email)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid email"})
		return
	}
	user.Password, _ = utils.HashPassword(user.Password)

	//Check if the user is already registered
	db, err := configs.ConnectDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error occurred, try again"})
		return
	}
	res := db.Where("email= ?", user.Email).First(&user)
	if res.RowsAffected > 0 {
		c.JSON(http.StatusForbidden, gin.H{"Message": "Email already registered", "Rows Affected": db.RowsAffected})
		return
	}
	registeredUser, err := services.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err})
		return
	}
	c.JSON(200, gin.H{"Message": "User Registered", "User": registeredUser})
}

func Login(c *gin.Context) {
	var req models.LoginRequest
	var user models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	db, err := configs.ConnectDb()
	if err != nil {
		return
	}
	if err := utils.CheckValidity(req.Email); !err {
		c.JSON(http.StatusUnauthorized, "Invalid email")
		return
	}

	res := db.Where("email= ?", req.Email).First(&user)
	if res.RowsAffected < 1 {
		c.JSON(404, "Email not found")
		return
	}

	isValid := utils.CompareHashAndPassword(user.Password, req.Password)
	if !isValid {
		c.JSON(http.StatusUnauthorized, "Invalid password")
		return
	}

	token, err := auth.GenerateToken(user.Name, time.Minute*15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error occurred"})
		return
	}

	newModel := &models.User{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		AccountId: user.AccountId,
		UpdatedAt: user.UpdatedAt,
	}

	RefreshToken, err := auth.GenerateToken(user.Name, time.Hour*24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Refresh token generation err")
	}

	//Access token payload
	atp, err := utils.GetPayload(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}

	//Refresh token payload
	rtp, err := utils.GetPayload(RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}

	userRes := &models.LoginResponse{
		AccessToken:           token,
		AccessTokenExpiresAt:  atp.ExpiresAt,
		RefreshToken:          RefreshToken,
		RefreshTokenExpiresAt: rtp.ExpiresAt,
		User:                  *newModel,
	}
	//Set refresh token in redis
	isSet, err := redis.SetRefreshToken(utils.UUIDString(rtp.TokenId), RefreshToken, time.Hour*24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message: ": "Error setting refresh token", "Error: ": err})
		return
	}
	if isSet {
		c.JSON(200, userRes)
		return
	}
	
}
