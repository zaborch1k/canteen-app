package usecase

import (
	"time"

	domAuth "canteen-app/internal/domain/auth"
	domUser "canteen-app/internal/domain/user"
)

type authUseCase struct {
	users       AuthRepository
	refreshRepo RefreshTokenRepository
	tokens      domAuth.TokenService
	accessTTL   time.Duration
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthUseCase(users AuthRepository, tokens domAuth.TokenService, accessTTL time.Duration, refreshRepo RefreshTokenRepository) *authUseCase {
	return &authUseCase{users: users, tokens: tokens, accessTTL: accessTTL, refreshRepo: refreshRepo}
}

func (uc *authUseCase) Register(login, password, name, surname, role string) (*Tokens, error) {
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

	access, err := uc.tokens.GenerateAccessToken(userID, user.Role)
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

func (uc *authUseCase) Login(login, password string) (*Tokens, error) {
	user, err := uc.users.GetUserByLogin(login)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := domUser.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, ErrInvalidCredentials
	}

	access, err := uc.tokens.GenerateAccessToken(user.ID, user.Role)
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

func (uc *authUseCase) GetUserByLogin(login string) (*domUser.User, error) {
	user, err := uc.users.GetUserByLogin(login)
	if err != nil {
		return &domUser.User{}, err
	}
	return user, nil
}

func (uc *authUseCase) Refresh(refreshToken string) (*Tokens, error) {
	userID, tokenID, err := uc.tokens.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, ErrInvalidRefresh
	}

	ok := uc.refreshRepo.IsValid(tokenID, userID)
	if !ok {
		return nil, ErrInvalidRefresh
	}

	uc.refreshRepo.Delete(tokenID)

	user, _ := uc.users.GetUserByID(userID)
	access, err := uc.tokens.GenerateAccessToken(userID, user.Role)
	if err != nil {
		return nil, err
	}

	newRefresh, newID, newExp, err := uc.tokens.GenerateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	uc.refreshRepo.Save(newID, userID, newExp)
	return &Tokens{AccessToken: access, RefreshToken: newRefresh}, nil
}
