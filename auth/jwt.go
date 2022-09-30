package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

const secretKeySize = 16

type Maker interface{
	CreateToken(username string, duration time.Duration)(string, error)
	VerifyToken(token string)(*Payload, error)
}

type JwtMaker struct{
	secretKey string
}

type Payload struct {
	TokenId   uuid.UUID `json:"tokenId"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiresAt time.Time `json:"expiredAt"`
}

//Generate new payload 
func NewPayload(username string, duration time.Duration) (*Payload){
	payload := &Payload{
		TokenId: uuid.NewV4(),
		Username: username,
		IssuedAt: time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload
}

//Add valid to payload 
func (p *Payload) Valid() error{
	if time.Now().After(p.ExpiresAt){
		return fmt.Errorf("token expired")
	}
	return nil
}

//Create new jwt maker
func NewMaker(secretKey string)(Maker, error){
	if len(secretKey)<secretKeySize{
		return nil, fmt.Errorf("secret key length too short. Required length: %v",secretKeySize)
	}
	return &JwtMaker{secretKey}, nil
}

//Create new jwt token
func (m *JwtMaker) CreateToken(username string, duration time.Duration)(string, error){
	payload := NewPayload(username, duration)
	jwtToken  := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	
	return jwtToken.SignedString([]byte(m.secretKey))
}

//Verify jwt token
func (m *JwtMaker) VerifyToken(token string)(*Payload, error){

	KeyFunc := func(token *jwt.Token)(interface{}, error){
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(m.secretKey), nil
	}

	jwtToken , err := jwt.ParseWithClaims(token, &Payload{}, KeyFunc)
	if err!=nil{
		return nil, fmt.Errorf("invalid token error: %v", err)
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok{
		return nil, fmt.Errorf("invalid token")
	}
	return payload, nil
}



