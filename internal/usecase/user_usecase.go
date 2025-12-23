package usecase

import "canteen-app/internal/domain"

type UserUseCase struct {
	users UserRepository
}

func NewUserUseCase(users UserRepository) *UserUseCase {
	return &UserUseCase{users: users}
}

func (uc *UserUseCase) RegisterUser(login, password, role string) {
	user := domain.User{
		Login:        login,
		PasswordHash: password,
		Role:         role,
	}

	uc.users.CreateUser(user)
}
