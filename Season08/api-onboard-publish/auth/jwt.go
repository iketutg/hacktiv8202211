package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTClaim struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

var key = []byte("12345678")

func ValidateJWT(token string) bool {
	tokenparse, err := jwt.ParseWithClaims(token, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		fmt.Println("Validate token error : ", err.Error())
		return false
	}

	claim, ok := tokenparse.Claims.(*JWTClaim)
	if !ok {
		fmt.Println("Parse Token error : ")
		return false
	}
	fmt.Println("claim Email : ", claim.Email)
	fmt.Println("claim UserName : ", claim.UserName)
	fmt.Println("claim ExpiresAt : ", claim.ExpiresAt)

	if claim.ExpiresAt < time.Now().Local().Unix() {
		fmt.Println("Token Expired ")
		return true
	}

	fmt.Println("Token Not Expired")
	return true
}
