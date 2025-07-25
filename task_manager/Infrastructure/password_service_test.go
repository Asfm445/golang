package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_Success(t *testing.T) {
	password := "mySecret123"

	hashed, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)
	assert.NotEqual(t, password, hashed)
}

func TestCheckPasswordHash_CorrectPassword(t *testing.T) {
	password := "securePassword!"
	hashed, _ := HashPassword(password)

	match := CheckPasswordHash(hashed, password)
	assert.True(t, match)
}

func TestCheckPasswordHash_WrongPassword(t *testing.T) {
	password := "correctPass"
	wrong := "wrongPass"
	hashed, _ := HashPassword(password)

	match := CheckPasswordHash(hashed, wrong)
	assert.False(t, match)
}
