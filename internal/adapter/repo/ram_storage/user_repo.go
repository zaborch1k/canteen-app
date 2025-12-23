package ram_storage

import (
	"errors"
	"fmt"

	"canteen-app/internal/domain"
)

type UserRepo struct {
	Users map[domain.UserID]domain.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		Users: make(map[domain.UserID]domain.User),
	}
}

func (ur *UserRepo) CreateUser(user domain.User) {
	ur.Users[user.ID] = user
	fmt.Print(ur.Users)
}

func (ur UserRepo) GetUserByID(id domain.UserID) (domain.User, error) {
	if user, ok := ur.Users[id]; ok {
		return user, nil
	}
	return domain.User{}, errors.New("user does not exist")
}
