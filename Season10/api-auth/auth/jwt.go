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

func GetToken(userName, email string) string {
	exp := time.Now().Add(3 * time.Minute)
	claim := JWTClaim{
		UserName:       userName,
		Email:          email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: exp.Unix()},
	}
	fmt.Println("Exp : ", exp.Unix())
	fmt.Println("Claim : ", claim)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedString, err := token.SignedString(key)
	fmt.Println("Error : ", err)
	fmt.Println("Signed String : ", signedString)
	return signedString
}
