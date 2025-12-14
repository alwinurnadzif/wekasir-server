package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = os.Getenv("jwt.secret_key")

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return webToken, nil
}

func VerifyToken(tokenJWT string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenJWT, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isValid := token.Claims.(jwt.MapClaims)
	if isValid && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token invalid")
}
