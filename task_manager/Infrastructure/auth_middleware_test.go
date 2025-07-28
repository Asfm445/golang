// file: auth_middleware_test.go
package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"task_manager/domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) VerifyToken(token string) (*domain.UserClaims, error) {
	args := m.Called(token)
	if claims := args.Get(0); claims != nil {
		return claims.(*domain.UserClaims), args.Error(1)
	}
	return nil, args.Error(1)
}

// Add this to satisfy the interface
func (m *MockTokenService) GenerateToken(userID, email, role string) (string, error) {
	args := m.Called(userID, email, role)
	return args.String(0), args.Error(1)
}

func performRequest(r http.Handler, method, path, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAuthMiddleware_MissingAuthorization(t *testing.T) {
	mockService := new(MockTokenService)
	router := gin.New()
	router.Use(AuthMiddleware(mockService, "admin"))
	router.GET("/secure", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := performRequest(router, "GET", "/secure", "")
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_InvalidFormat(t *testing.T) {
	mockService := new(MockTokenService)
	router := gin.New()
	router.Use(AuthMiddleware(mockService, "admin"))
	router.GET("/secure", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := performRequest(router, "GET", "/secure", "InvalidTokenHere")
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_TokenVerificationFails(t *testing.T) {
	mockService := new(MockTokenService)
	mockService.On("VerifyToken", "fakeToken").Return(nil, domain.ErrTokenInvalid)

	router := gin.New()
	router.Use(AuthMiddleware(mockService, "admin"))
	router.GET("/secure", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := performRequest(router, "GET", "/secure", "Bearer fakeToken")
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_InsufficientPermissions(t *testing.T) {
	mockService := new(MockTokenService)
	claims := &domain.UserClaims{
		UserID: "1", Email: "user@example.com", Role: "user",
	}
	mockService.On("VerifyToken", "validUserToken").Return(claims, nil)

	router := gin.New()
	router.Use(AuthMiddleware(mockService, "admin"))
	router.GET("/secure", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := performRequest(router, "GET", "/secure", "Bearer validUserToken")
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestAuthMiddleware_SuccessfulAuthorization(t *testing.T) {
	mockService := new(MockTokenService)
	claims := &domain.UserClaims{
		UserID: "1", Email: "admin@example.com", Role: "admin",
	}
	mockService.On("VerifyToken", "adminToken").Return(claims, nil)

	router := gin.New()
	router.Use(AuthMiddleware(mockService, "admin"))
	router.GET("/secure", func(c *gin.Context) {
		userEmail, _ := c.Get("user_email")
		c.JSON(http.StatusOK, gin.H{"email": userEmail})
	})

	w := performRequest(router, "GET", "/secure", "Bearer adminToken")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "admin@example.com")
}
