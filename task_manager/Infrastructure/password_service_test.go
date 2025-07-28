package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	hasher := BcryptHasher{}
	password := "securePassword123"

	hashed, err := hasher.HashPassword(password)

	assert.NoError(t, err, "Hashing should not produce an error")
	assert.NotEmpty(t, hashed, "Hashed password should not be empty")
	assert.NotEqual(t, password, hashed, "Hashed password should not be equal to original")
}

func TestCheckPassword_CorrectPassword(t *testing.T) {
	hasher := BcryptHasher{}
	password := "correctPassword"

	hashed, _ := hasher.HashPassword(password)
	match := hasher.CheckPassword(hashed, password)

	assert.True(t, match, "CheckPassword should return true for correct password")
}

func TestCheckPassword_IncorrectPassword(t *testing.T) {
	hasher := BcryptHasher{}
	password := "correctPassword"
	wrongPassword := "wrongPassword"

	hashed, _ := hasher.HashPassword(password)
	match := hasher.CheckPassword(hashed, wrongPassword)

	assert.False(t, match, "CheckPassword should return false for incorrect password")
}
