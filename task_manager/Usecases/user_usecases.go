package usecases

import "task_manager/domain"

type UserUseCase struct {
	Repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{Repo: r}
}

func (u *UserUseCase) Register(user domain.User) error {
	return u.Repo.Register(user)
}

func (u *UserUseCase) Login(email, password string) (string, error) {
	return u.Repo.Login(email, password)
}

func (u *UserUseCase) Promote(email string) error {
	return u.Repo.Promote(email)
}
