package usecase

import (
	"canteen-app/internal/domain"
)

type UserUseCase struct {
	users UserRepository
}

func NewUserUseCase(users UserRepository) *UserUseCase {
	return &UserUseCase{users: users}
}

func (uc *UserUseCase) RegisterUser(login, password, name, surname, role string) {
	user := domain.User{
		Login:        login,
		PasswordHash: password,
		Name:         name,
		Surname:      surname,
		Role:         role,
	}

	uc.users.CreateUser(user)
}

func (uc *UserUseCase) GetUserByLogin(login string) (*domain.User, error) {
	user, err := uc.users.GetUserByLogin(login)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}
