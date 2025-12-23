package usecase

import "canteen-app/internal/domain"

type UserRepository interface {
	CreateUser(user domain.User)
	GetUserByID(id domain.UserID) (domain.User, error)
}
