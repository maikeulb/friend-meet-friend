package auth

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SIGN_KEY = []byte("s3cr3t") // move to environmental variable

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SIGN_KEY, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func GenerateToken(user *User) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.Email,
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 5).Unix(),
	})
	fmt.Println(reflect.TypeOf(user.ID))
	tokenString, err := token.SignedString(SIGN_KEY)
	if err != nil {
		return err
	}
	user.Token = tokenString

	return nil
}
