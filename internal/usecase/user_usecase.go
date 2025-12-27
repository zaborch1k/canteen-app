package usecase

import (
	domAuth "canteen-app/internal/domain/auth"
	domUser "canteen-app/internal/domain/user"
	"time"
)

type userUseCase struct {
	users       UserRepository
	refreshRepo RefreshTokenRepository
	tokens      domAuth.TokenService
	accessTTL   time.Duration
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewUserUseCase(users UserRepository, tokens domAuth.TokenService, accessTTL time.Duration, refreshRepo RefreshTokenRepository) *userUseCase {
	return &userUseCase{users: users, tokens: tokens, accessTTL: accessTTL, refreshRepo: refreshRepo}
}

func (uc *userUseCase) Register(login, password, name, surname, role string) (*Tokens, error) {
	if _, err := uc.users.GetUserByLogin(login); err == nil {
		return nil, ErrUserExists
	}

	hash, err := domUser.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := domUser.User{
		Login:        login,
		PasswordHash: hash,
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
		return nil, err
	}

	refresh, refreshID, refreshExp, err := uc.tokens.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	uc.refreshRepo.Save(refreshID, user.ID, refreshExp)

	return &Tokens{AccessToken: access, RefreshToken: refresh}, nil
}

func (uc *userUseCase) Login(login, password string) (*Tokens, error) {
	user, err := uc.users.GetUserByLogin(login)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := domUser.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, ErrInvalidCredentials
	}

	claims := domAuth.Claims{
		UserID:    user.ID,
		Role:      user.Role,
		ExpiresAt: time.Now().Add(uc.accessTTL),
	}

	access, err := uc.tokens.GenerateAccesToken(claims)
	if err != nil {
		return nil, err
	}

	refresh, refreshID, refreshExp, err := uc.tokens.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	uc.refreshRepo.Save(refreshID, user.ID, refreshExp)

	return &Tokens{AccessToken: access, RefreshToken: refresh}, nil
}

func (uc *userUseCase) GetUserByLogin(login string) (*domUser.User, error) {
	user, err := uc.users.GetUserByLogin(login)
	if err != nil {
		return &domUser.User{}, err
	}
	return user, nil
}
