package usecases

import "task_manager/domain"

type UserUseCaseInterface interface {
	Register(user domain.User) error
	Login(email, password string) (string, error)
	Promote(email string) error
}

type UserUseCase struct {
	Repo   domain.UserRepository
	Hasher domain.Hasher
	Token  domain.TokenService
}

func NewUserUseCase(r domain.UserRepository, h domain.Hasher, t domain.TokenService) *UserUseCase {
	return &UserUseCase{
		Repo:   r,
		Hasher: h,
		Token:  t,
	}
}

func (u *UserUseCase) Register(user domain.User) error {
	hashed, err := u.Hasher.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	return u.Repo.Register(user)
}

func (u *UserUseCase) Login(email, password string) (string, error) {
	user, err := u.Repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if !u.Hasher.CheckPassword(user.Password, password) {
		return "", domain.ErrInvalidCredentials
	}

	return u.Token.GenerateToken(user.ID, user.Email, user.Role)
}

func (u *UserUseCase) Promote(email string) error {
	return u.Repo.Promote(email)
}
