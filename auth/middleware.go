package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const authHeaderKey="authorization"
const authorizationType ="bearer"
const tokenPayload="token_payload"

func AuthMiddleware(maker Maker) gin.HandlerFunc{
	return func (c *gin.Context){

		authHeader := c.GetHeader(authHeaderKey)
		if len(authHeader)==0{
			err := errors.New("authorization header not provided")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error: ": err, "Message: ":"Header not provided"})
		}

		fields := strings.Fields(authHeader)
		if len(authHeader)<2{
			err := errors.New("invalid header provided")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error: ": err, "Message: ":"Invalid headers"})
		}

		authType := strings.ToLower(fields[0])
		if authType != authorizationType{
			err := errors.New("invalid authorization type provided")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error: ": err, "Message: ":"Invalid authorization type", "Type: ":authType})
		}

		accessToken := fields[1]
		payload, err := maker.VerifyToken(accessToken)
		if err != nil{
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error: ": err, "Message: ":"Invalid headers"})
		}
		
		c.Set(tokenPayload, payload)
		c.Next()
	}
}