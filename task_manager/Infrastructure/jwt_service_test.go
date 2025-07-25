// Infrastructure/jwt_service_test.go
package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidateToken(t *testing.T) {
	userID := "1"
	email := "test@example.com"
	role := "admin"

	// Generate token
	tokenStr, err := GenerateToken(userID, email, role)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

}
