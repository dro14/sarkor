package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Login  string
	UserID uint
}

var jwtSecret = []byte("sarkor")

func GenerateToken(claims *Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":   claims.Login,
		"user_id": claims.UserID,
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		login := claims["login"].(string)
		userID := uint(claims["user_id"].(float64))

		return &Claims{login, userID}, nil
	}

	return nil, errors.New("invalid token")
}
