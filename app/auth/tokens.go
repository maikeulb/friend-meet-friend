package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SIGN_KEY = []byte("s3cr3t")

type CustomClaims struct {
	sub string
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(sub string) (string, error)
}

func ParseToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SIGN_KEY, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func GenerateToken(user User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 24).Unix()

	claims := CustomClaims{
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "go.micro.srv.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SIGN_KEY)
}
