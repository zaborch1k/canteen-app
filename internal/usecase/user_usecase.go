package usecase

import (
	domAuth "canteen-app/internal/domain/auth"
	domUser "canteen-app/internal/domain/user"
	"time"
)

type userUseCase struct {
	users     UserRepository
	tokens    domAuth.TokenService
	accessTTL time.Duration
}

func NewUserUseCase(users UserRepository, tokens domAuth.TokenService, accessTTL time.Duration) *userUseCase {
	return &userUseCase{users: users, tokens: tokens, accessTTL: accessTTL}
}

func (uc *userUseCase) RegisterUser(login, password, name, surname, role string) (accessToken string, err error) {
	user := domUser.User{
		Login:        login,
		PasswordHash: password,
		Name:         name,
		Surname:      surname,
		Role:         role,
	}

	userID := uc.users.CreateUser(user)

	claims := domAuth.Claims{
		UserID:    userID,
		Role:      user.Role,
		ExpiresAt: time.Now().Add(uc.accessTTL),
	}

	access, err := uc.tokens.GenerateAccesToken(claims)
	if err != nil {
		return "", err
	}

	return access, nil
}

func (uc *userUseCase) GetUserByLogin(login string) (*domUser.User, error) {
	user, err := uc.users.GetUserByLogin(login)
	if err != nil {
		return &domUser.User{}, err
	}
	return user, nil
}
