package usecases

import (
	"task_manager/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mocks ---

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Register(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) FindByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepo) Promote(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

type MockHasher struct {
	mock.Mock
}

func (m *MockHasher) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockHasher) CheckPassword(hashed, password string) bool {
	args := m.Called(hashed, password)
	return args.Bool(0)
}

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) GenerateToken(userID, email, role string) (string, error) {
	args := m.Called(userID, email, role)
	return args.String(0), args.Error(1)
}

func (m *MockTokenService) VerifyToken(tokenStr string) (*domain.UserClaims, error) {
	args := m.Called(tokenStr)
	return args.Get(0).(*domain.UserClaims), args.Error(1)
}

func TestUserUseCase_Login_Success(t *testing.T) {
	repo := new(MockUserRepo)
	hasher := new(MockHasher)
	token := new(MockTokenService)

	user := domain.User{
		ID:       "user123",
		Email:    "test@example.com",
		Password: "hashed-pass",
		Role:     "user",
	}

	repo.On("FindByEmail", "test@example.com").Return(user, nil)
	hasher.On("CheckPassword", "hashed-pass", "plain-pass").Return(true)
	token.On("GenerateToken", "user123", "test@example.com", "user").Return("fake-jwt", nil)

	usecase := NewUserUseCase(repo, hasher, token)

	result, err := usecase.Login("test@example.com", "plain-pass")

	assert.NoError(t, err)
	assert.Equal(t, "fake-jwt", result)

	repo.AssertExpectations(t)
	hasher.AssertExpectations(t)
	token.AssertExpectations(t)
}

func TestUserUseCase_Login_InvalidPassword(t *testing.T) {
	repo := new(MockUserRepo)
	hasher := new(MockHasher)
	token := new(MockTokenService)

	user := domain.User{
		ID:       "user123",
		Email:    "test@example.com",
		Password: "hashed-pass",
		Role:     "user",
	}

	repo.On("FindByEmail", "test@example.com").Return(user, nil)
	hasher.On("CheckPassword", "hashed-pass", "wrong-pass").Return(false)

	usecase := NewUserUseCase(repo, hasher, token)

	result, err := usecase.Login("test@example.com", "wrong-pass")

	assert.ErrorIs(t, err, domain.ErrInvalidCredentials)
	assert.Empty(t, result)

	repo.AssertExpectations(t)
	hasher.AssertExpectations(t)
}

func TestUserUseCase_Register_Success(t *testing.T) {
	repo := new(MockUserRepo)
	hasher := new(MockHasher)
	token := new(MockTokenService)

	user := domain.User{
		Email:    "test@example.com",
		Password: "plain-pass",
	}

	hasher.On("HashPassword", "plain-pass").Return("hashed-pass", nil)
	expectedUser := user
	expectedUser.Password = "hashed-pass"

	repo.On("Register", expectedUser).Return(nil)

	usecase := NewUserUseCase(repo, hasher, token)
	err := usecase.Register(user)

	assert.NoError(t, err)
	hasher.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestUserUseCase_Register_HashFails(t *testing.T) {
	repo := new(MockUserRepo)
	hasher := new(MockHasher)
	token := new(MockTokenService)

	user := domain.User{
		Email:    "test@example.com",
		Password: "plain-pass",
	}

	hasher.On("HashPassword", "plain-pass").Return("", assert.AnError)

	usecase := NewUserUseCase(repo, hasher, token)
	err := usecase.Register(user)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	repo.AssertNotCalled(t, "Register", mock.Anything)
	hasher.AssertExpectations(t)
}

func TestUserUseCase_Register_RepoFails(t *testing.T) {
	repo := new(MockUserRepo)
	hasher := new(MockHasher)
	token := new(MockTokenService)

	user := domain.User{
		Email:    "test@example.com",
		Password: "plain-pass",
	}

	hasher.On("HashPassword", "plain-pass").Return("hashed-pass", nil)

	expectedUser := user
	expectedUser.Password = "hashed-pass"

	repo.On("Register", expectedUser).Return(assert.AnError)

	usecase := NewUserUseCase(repo, hasher, token)
	err := usecase.Register(user)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	hasher.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestUserUseCase_Promote_Success(t *testing.T) {
	repo := new(MockUserRepo)
	hasher := new(MockHasher)
	token := new(MockTokenService)

	repo.On("Promote", "user@example.com").Return(nil)

	usecase := NewUserUseCase(repo, hasher, token)
	err := usecase.Promote("user@example.com")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUserUseCase_Promote_Fails(t *testing.T) {
	repo := new(MockUserRepo)
	hasher := new(MockHasher)
	token := new(MockTokenService)

	repo.On("Promote", "user@example.com").Return(assert.AnError)

	usecase := NewUserUseCase(repo, hasher, token)
	err := usecase.Promote("user@example.com")

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	repo.AssertExpectations(t)
}
