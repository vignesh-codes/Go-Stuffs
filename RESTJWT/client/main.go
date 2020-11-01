package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT() (string, string) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignalString(mySigningKey)

	if err != nil {
		fmt.Println("Errorrr", err.Error())
		return "", err
	}

}
func main() {
	fmt.Println("ESimple Client")
	tokenString, err := GenerateJWT()

	if err != nil {
		fmt.Println("Error Generating the String")
	




	}

	fmt.Println(tokenString)
}
