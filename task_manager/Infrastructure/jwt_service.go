package infrastructure

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte("your_jwt_secret")

func GenerateToken(userID, email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
		"iat":     time.Now().Unix(),
	})
	return token.SignedString(JwtSecret)
}
