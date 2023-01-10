package main

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

func main() {
	exp := time.Now().Add(3 * time.Minute)
	claim := JWTClaim{
		UserName:       "iketutg",
		Email:          "iketutg@gmail.com",
		StandardClaims: jwt.StandardClaims{ExpiresAt: exp.Unix()},
	}
	fmt.Println("Exp : ", exp.Unix())
	fmt.Println("Claim : ", claim)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedString, err := token.SignedString(key)
	fmt.Println("Error : ", err)
	fmt.Println("Signed String : ", signedString)

	fmt.Println("============")
	tokenexpired := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImlrZXR1dGciLCJlbWFpbCI6ImlrZXR1dGdAZ21haWwuY29tIiwiZXhwIjoxNjcyNzU0NDMwfQ.adCWfkxzrLj1WGnmkGPHxWj7MGRN8K0qGXfPTJXnjW0"

	fmt.Println("Validate old")
	ValidateJWT(tokenexpired)

	fmt.Println("============")
	fmt.Println("Validate New ")
	ValidateJWT(signedString)

}

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
