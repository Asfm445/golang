package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task_manager/domain"
	"testing"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUsecase) Login(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) Promote(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRegister_Success(t *testing.T) {
	mockUC := new(MockUserUsecase)
	ctrl := NewUserController(mockUC)

	user := domain.User{Email: "user@example.com", Password: "1234"}
	mockUC.On("Register", user).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", ctrl.Register)

	body, _ := json.Marshal(user)
	resp := performRequest(router, "POST", "/register", body)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockUC.AssertExpectations(t)
}

func TestRegister_InvalidJSON(t *testing.T) {
	ctrl := NewUserController(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", ctrl.Register)

	resp := performRequest(router, "POST", "/register", []byte(`invalid-json`))

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestRegister_Failure(t *testing.T) {
	mockUC := new(MockUserUsecase)
	ctrl := NewUserController(mockUC)

	user := domain.User{Email: "user@example.com", Password: "1234"}
	mockUC.On("Register", user).Return(assert.AnError)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", ctrl.Register)

	body, _ := json.Marshal(user)
	resp := performRequest(router, "POST", "/register", body)

	assert.Equal(t, http.StatusConflict, resp.Code)
	mockUC.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockUC := new(MockUserUsecase)
	ctrl := NewUserController(mockUC)

	input := map[string]string{"email": "test@example.com", "password": "1234"}
	mockUC.On("Login", input["email"], input["password"]).Return("mockedToken", nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", ctrl.Login)

	body, _ := json.Marshal(input)
	resp := performRequest(router, "POST", "/login", body)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUC.AssertExpectations(t)
}

func TestLogin_InvalidJSON(t *testing.T) {
	ctrl := NewUserController(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", ctrl.Login)

	resp := performRequest(router, "POST", "/login", []byte(`bad-json`))

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestLogin_InvalidCreds(t *testing.T) {
	mockUC := new(MockUserUsecase)
	ctrl := NewUserController(mockUC)

	input := map[string]string{"email": "test@example.com", "password": "wrongpass"}
	mockUC.On("Login", input["email"], input["password"]).Return("", domain.ErrInvalidCredentials)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", ctrl.Login)

	body, _ := json.Marshal(input)
	resp := performRequest(router, "POST", "/login", body)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	mockUC.AssertExpectations(t)
}

func TestPromote_Success(t *testing.T) {
	mockUC := new(MockUserUsecase)
	ctrl := NewUserController(mockUC)

	input := map[string]string{"email": "test@example.com"}
	mockUC.On("Promote", input["email"]).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PATCH("/promote", ctrl.Promote)

	body, _ := json.Marshal(input)
	resp := performRequest(router, "PATCH", "/promote", body)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUC.AssertExpectations(t)
}

func TestPromote_InvalidJSON(t *testing.T) {
	ctrl := NewUserController(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PATCH("/promote", ctrl.Promote)

	resp := performRequest(router, "PATCH", "/promote", []byte(`bad-json`))

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestPromote_Failure(t *testing.T) {
	mockUC := new(MockUserUsecase)
	ctrl := NewUserController(mockUC)

	input := map[string]string{"email": "test@example.com"}
	mockUC.On("Promote", input["email"]).Return(assert.AnError)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PATCH("/promote", ctrl.Promote)

	body, _ := json.Marshal(input)
	resp := performRequest(router, "PATCH", "/promote", body)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockUC.AssertExpectations(t)
}
