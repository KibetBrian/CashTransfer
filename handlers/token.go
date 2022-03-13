package handlers

import (
	"net/http"
	"time"

	"github.com/KibetBrian/fisa/auth"
	"github.com/KibetBrian/fisa/database/redis"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
)


func RefreshToken(c *gin.Context){
	var RefreshToken models.RefreshTokenReq
	c.ShouldBindJSON(&RefreshToken)

	p, err := utils.GetPayload(RefreshToken.RefreshToken)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, "")
	}	

	//Check db if the token exists in redis
	s, err := redis.GetRefreshToken(p.Username)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message ": "Token not found","Token: ":s})
	}

	sk, err := utils.GetEnvVal("JWT_SECRET_KEY")
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error")
	}

	maker, err := auth.NewMaker(sk)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error")
	}

	//Generate new token
	token, err := maker.CreateToken(p.Username, time.Hour*24)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error generating token")
	}
	
	//Set new token generated
	isSet, err := redis.SetRefreshToken(utils.UUIDString(p.TokenId), token, time.Hour*24)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Error setting token", "Error: ": err})
	}
	
	if isSet {
		c.JSON(http.StatusAccepted, "Token set")
		return
	}
}