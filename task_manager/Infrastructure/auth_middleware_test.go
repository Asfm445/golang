// Add at the top of the test file
package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testSecret = []byte("test_secret") // override JwtSecret in tests

func generateTestToken(userID, email string, role string, expired bool) string {
	exp := time.Now().Add(time.Hour * 1).Unix()
	if expired {
		exp = time.Now().Add(-time.Hour).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     exp,
		"iat":     time.Now().Unix(),
	})

	tokenStr, _ := token.SignedString([]byte(testSecret))
	return tokenStr
}

func init() {
	JwtSecret = testSecret
}

func TestAuthMiddleware_NoToken(t *testing.T) {
	router := gin.New()
	router.GET("/protected", AuthMiddleware("admin"), func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
}

func TestAuthMiddleware_InvalidFormat(t *testing.T) {
	router := gin.New()
	router.GET("/protected", AuthMiddleware("admin"), func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "invalidTokenHere")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
}

func TestAuthMiddleware_ExpiredToken(t *testing.T) {
	router := gin.New()
	router.GET("/protected", AuthMiddleware("admin"), func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	token := generateTestToken("1", "awel@example.com", "admin", true)
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
}

func TestAuthMiddleware_ValidAdminToken(t *testing.T) {
	router := gin.New()
	router.GET("/protected", AuthMiddleware("admin"), func(c *gin.Context) {
		email := c.MustGet("user_email").(string)
		c.JSON(200, gin.H{"email": email})
	})

	token := generateTestToken("1", "admin@example.com", "admin", false)
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "admin@example.com")
}

func TestAuthMiddleware_InsufficientRole(t *testing.T) {
	router := gin.New()
	router.GET("/protected", AuthMiddleware("admin"), func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	token := generateTestToken("1", "user@example.com", "user", false)
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 403, resp.Code)
}
