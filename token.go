package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

var (
	tokenInitialized = false
	tokenSecret      = []byte("")
)

var (
	ErrTokenNotInitialized = fmt.Errorf("util-go/token.go - token not initialized")
)

func Generate(userID interface{}) (string, error) {
	if !isTokenInitialized() {
		return "", ErrTokenNotInitialized
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID,
	})

	return token.SignedString(tokenSecret)
}

func Validate(tokenString string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tokenSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, false
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, true
}

// MARK: Utilities

func SetTokenSecret(TokenSecret string) {
	tokenSecret = []byte(TokenSecret)
	tokenInitialized = true
}
func isTokenInitialized() bool {
	return tokenInitialized
}
