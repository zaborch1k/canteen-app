package usecase

import (
	domUser "canteen-app/internal/domain/user"
)

type userUseCase struct {
	users UserRepository
}

func NewUserUseCase(users UserRepository) *userUseCase {
	return &userUseCase{users: users}
}

func (uc *userUseCase) RegisterUser(login, password, name, surname, role string) {
	user := domUser.User{
		Login:        login,
		PasswordHash: password,
		Name:         name,
		Surname:      surname,
		Role:         role,
	}

	uc.users.CreateUser(user)
}

func (uc *userUseCase) GetUserByLogin(login string) (*domUser.User, error) {
	user, err := uc.users.GetUserByLogin(login)
	if err != nil {
		return &domUser.User{}, err
	}
	return user, nil
}
