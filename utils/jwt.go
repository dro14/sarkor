package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// Claims is a struct that will be encoded to a JWT.
type Claims struct {
	Login  string
	UserID uint
}

// jwtSecret is a secret key that will be used to create a signature.
var jwtSecret = []byte("sarkor")

// GenerateToken generates a jwt token and assign a login and user_id to its claims and return it.
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

// ValidateToken validates the jwt token and return its claims if it's valid.
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
