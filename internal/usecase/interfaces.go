package usecase

import "canteen-app/internal/domain"

type UserRepository interface {
	CreateUser(user domain.User)
	GetUserByID(id domain.UserID) (*domain.User, error)
	GetUserByLogin(login string) (*domain.User, error)
}

type UserUseCase interface {
	RegisterUser(login, password, name, surname, role string)
	GetUserByLogin(login string) (*domain.User, error)
}
