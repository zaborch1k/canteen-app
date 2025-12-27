package usecase

import domUser "canteen-app/internal/domain/user"

type UserRepository interface {
	CreateUser(user domUser.User) domUser.UserID
	GetUserByID(id domUser.UserID) (*domUser.User, error)
	GetUserByLogin(login string) (*domUser.User, error)
}

type UserUseCase interface {
	Register(login, password, name, surname, role string) (accessToken string, err error)
	Login(login, password string) (accessToken string, err error)
	GetUserByLogin(login string) (*domUser.User, error)
}
