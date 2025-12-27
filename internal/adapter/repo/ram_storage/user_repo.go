package ram_storage

import (
	"errors"
	"math/rand"
	"time"

	domUser "canteen-app/internal/domain/user"
	"canteen-app/internal/usecase"
)

type UserRepo struct {
	Users map[domUser.UserID]domUser.User
}

var _ usecase.UserRepository = (*UserRepo)(nil)

func NewUserRepo() *UserRepo {
	return &UserRepo{
		Users: make(map[domUser.UserID]domUser.User),
	}
}

func (ur *UserRepo) CreateUser(user domUser.User) domUser.UserID {
	rand.Seed(time.Now().UnixNano())
	user.ID = domUser.UserID(rand.Int63n(int64(234)))
	ur.Users[user.ID] = user
	return user.ID
}

func (ur UserRepo) GetUserByID(id domUser.UserID) (*domUser.User, error) {
	if user, ok := ur.Users[id]; ok {
		return &user, nil
	}
	return &domUser.User{}, errors.New("user does not exist")
}

func (uc UserRepo) GetUserByLogin(login string) (*domUser.User, error) {
	for _, val := range uc.Users {
		if val.Login == login {
			return &val, nil
		}
	}
	return &domUser.User{}, errors.New("user does not exist")
}
