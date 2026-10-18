package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("1")

func GenerateToken(UserID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		claims,
	)

	return token.SignedString(secretKey)
}
