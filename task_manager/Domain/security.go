package domain

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(hashed, password string) bool
}

type UserClaims struct {
	UserID string
	Email  string
	Role   string
}

type TokenService interface {
	GenerateToken(userID, email, role string) (string, error)
	VerifyToken(tokenStr string) (*UserClaims, error)
}
