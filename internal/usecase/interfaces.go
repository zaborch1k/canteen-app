package usecase

import domUser "canteen-app/internal/domain/user"

type UserRepository interface {
	CreateUser(user domUser.User) domUser.UserID
	GetUserByID(id domUser.UserID) (*domUser.User, error)
	GetUserByLogin(login string) (*domUser.User, error)
}

type UserUseCase interface {
	RegisterUser(login, password, name, surname, role string)
	GetUserByLogin(login string) (*domUser.User, error)
}
