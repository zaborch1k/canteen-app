package ram_storage

import (
	"errors"

	"canteen-app/internal/domain"
	"canteen-app/internal/usecase"
)

type UserRepo struct {
	Users map[domain.UserID]domain.User
}

var _ usecase.UserRepository = (*UserRepo)(nil)

func NewUserRepo() *UserRepo {
	return &UserRepo{
		Users: make(map[domain.UserID]domain.User),
	}
}

func (ur *UserRepo) CreateUser(user domain.User) {
	ur.Users[user.ID] = user
}

func (ur UserRepo) GetUserByID(id domain.UserID) (domain.User, error) {
	if user, ok := ur.Users[id]; ok {
		return user, nil
	}
	return domain.User{}, errors.New("user does not exist")
}
