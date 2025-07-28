package infrastructure

import (
	"task_manager/domain"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAndVerifyToken(t *testing.T) {
	jwtService := NewJWTService()

	userID := "123"
	email := "test@example.com"
	role := "admin"

	token, err := jwtService.GenerateToken(userID, email, role)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := jwtService.VerifyToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)
	assert.Equal(t, role, claims.Role)
}

func TestVerifyToken_InvalidToken(t *testing.T) {
	jwtService := NewJWTService()
	invalidToken := "this.is.not.valid"

	claims, err := jwtService.VerifyToken(invalidToken)

	assert.Nil(t, claims)
	assert.Equal(t, domain.ErrTokenInvalid, err)
}

func TestVerifyToken_ExpiredToken(t *testing.T) {
	jwtService := NewJWTService()

	// Create an expired token manually
	expiredClaims := jwt.MapClaims{
		"user_id": "expired123",
		"email":   "expired@example.com",
		"role":    "user",
		"exp":     time.Now().Add(-time.Hour).Unix(), // already expired
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	signedToken, _ := token.SignedString([]byte(jwtService.SecretKey))

	claims, err := jwtService.VerifyToken(signedToken)

	assert.Nil(t, claims)
	assert.Equal(t, domain.ErrTokenInvalid, err)
}
