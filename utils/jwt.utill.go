package utils

import "github.com/golang-jwt/jwt/v5"

var secret_key = "SECRET_KEY"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webToken, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}

	return webToken, nil
}
