package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/KibetBrian/fisa/auth"
	"github.com/KibetBrian/fisa/database/redis"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	var RefreshToken models.RefreshTokenReq
	c.ShouldBindJSON(&RefreshToken)

	p, err := utils.GetPayload(RefreshToken.RefreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "")
	}

	//Check db if the token exists in redis
	s, err := redis.GetRefreshToken(p.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message ": "Token not found", "Token: ": s})
	}

	sk, isFound:=os.LookupEnv("JWT_SECRET_KEY")
	if !isFound {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error")
	}

	maker, err := auth.NewMaker(sk)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error")
	}

	//Generate new access token
	accessToken, err := maker.CreateToken(p.Username, time.Hour*24)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error generating access token")
	}
	ap, err := utils.GetPayload(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error getting payload")
	}

	//Generate new refresh token
	refreshToken, err := maker.CreateToken(p.Username, time.Hour*24)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error generating refreshtoken")
	}
	rp, err := utils.GetPayload(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error getting payload")
	}

	//Construct response request
	rtr := &models.RefreshTokenResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  ap.ExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: rp.ExpiresAt,
	}

	//Set new token generated
	isSet, err := redis.SetRefreshToken(utils.UUIDString(p.TokenId), refreshToken, time.Hour*24)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Error setting token", "Error: ": err})
	}

	if isSet {
		c.JSON(http.StatusAccepted, rtr)
		return
	}
}
