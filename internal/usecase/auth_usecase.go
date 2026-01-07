package usecase

import (
	domAuth "canteen-app/internal/domain/auth"
	domUser "canteen-app/internal/domain/user"
)

type authUseCase struct {
	users       UserRepository
	refreshRepo RefreshTokenRepository
	tokens      TokenService
	hasher      PasswordHasher
}

func NewAuthUseCase(users UserRepository, tokens TokenService, refreshRepo RefreshTokenRepository, hasher PasswordHasher) *authUseCase {
	return &authUseCase{users: users, tokens: tokens, refreshRepo: refreshRepo, hasher: hasher}
}

func (uc *authUseCase) Register(login, password, name, surname, role string) (*domAuth.Tokens, error) {
	if _, err := uc.users.GetUserByLogin(login); err == nil {
		return nil, ErrUserExists
	}

	hash, err := uc.hasher.Hash(password)
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

	return &domAuth.Tokens{AccessToken: access, RefreshToken: refresh}, nil
}

func (uc *authUseCase) Login(login, password string) (*domAuth.Tokens, error) {
	user, err := uc.users.GetUserByLogin(login)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := uc.hasher.Compare(user.PasswordHash, password); err != nil {
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

	return &domAuth.Tokens{AccessToken: access, RefreshToken: refresh}, nil
}

func (uc *authUseCase) GetUserByLogin(login string) (*domUser.User, error) {
	user, err := uc.users.GetUserByLogin(login)
	if err != nil {
		return &domUser.User{}, err
	}
	return user, nil
}

func (uc *authUseCase) Refresh(refreshToken string) (*domAuth.Tokens, error) {
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
	return &domAuth.Tokens{AccessToken: access, RefreshToken: newRefresh}, nil
}

func (uc *authUseCase) RevokeRefreshToken(refreshToken string) error {
	userID, tokenID, err := uc.tokens.ParseRefreshToken(refreshToken)
	if err != nil {
		return ErrInvalidRefresh
	}

	ok := uc.refreshRepo.IsValid(tokenID, userID)
	if !ok {
		return ErrInvalidRefresh
	}

	uc.refreshRepo.Delete(tokenID)
	return nil
}
