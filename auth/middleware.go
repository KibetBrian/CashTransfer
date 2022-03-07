package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const authHeaderKey="authorization"
const authorizationType ="Bearer"
const tokenPayload="token_payload"

func AuthMiddleware(maker Maker) gin.HandlerFunc{
	return func (c *gin.Context){

		authHeader := c.GetHeader(authHeaderKey)
		if len(authHeader)==0{
			err := errors.New("authorization header not provided")
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}

		fields := strings.Fields(authHeader)
		if len(authHeader)<2{
			err := errors.New("invalid header provided")
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}

		authType := strings.ToLower(fields[0])
		if authType != authorizationType{
			err := errors.New("invalid header provided")
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}

		accessToken := fields[1]
		payload, err := maker.VerifyToken(accessToken)
		if err != nil{
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
		}
		
		c.Set(tokenPayload, payload)
		c.Next()
	}
}