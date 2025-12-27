package jwtadapter

import (
	"fmt"
	"time"

	domAuth "canteen-app/internal/domain/auth"
	domUser "canteen-app/internal/domain/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTTokenService struct {
	accessSecret  []byte
	refreshSecret []byte
	accessTTL     time.Duration
	refreshTTL    time.Duration
	issuer        string
}

func NewJWTTokenService(accessSecret, refreshSecret []byte, accessTTL, refreshTTL time.Duration, issuer string) *JWTTokenService {
	return &JWTTokenService{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessTTL:     accessTTL,
		refreshTTL:    refreshTTL,
		issuer:        issuer,
	}
}

type accessClaims struct {
	UserID domUser.UserID `json:"sub"`
	Role   string         `json:"role"`
	jwt.RegisteredClaims
}

type refreshClaims struct {
	UserID  domUser.UserID `json:"sub"`
	TokenID string         `json:"tid"`
	jwt.RegisteredClaims
}

func (s *JWTTokenService) GenerateAccessToken(userID domUser.UserID, role string) (string, error) {
	exp := time.Now().Add(s.accessTTL)
	claims := accessClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.accessSecret)
}

func (s *JWTTokenService) ParseAccessToken(tokenStr string) (domAuth.Claims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &accessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.accessSecret, nil
	})

	if err != nil || !t.Valid {
		return domAuth.Claims{}, fmt.Errorf("invalid token: %w", err)
	}

	cl, ok := t.Claims.(*accessClaims)
	if !ok {
		return domAuth.Claims{}, fmt.Errorf("invalid claims type")
	}

	return domAuth.Claims{
		UserID:    cl.UserID,
		Role:      cl.Role,
		ExpiresAt: cl.ExpiresAt.Time,
	}, nil
}

func (s *JWTTokenService) GenerateRefreshToken(userID domUser.UserID) (string, string, time.Time, error) {
	exp := time.Now().Add(s.refreshTTL)
	id := uuid.NewString()
	claims := refreshClaims{
		UserID:  userID,
		TokenID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			ID:        id,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := t.SignedString(s.refreshSecret)
	return str, id, exp, err
}

func (s *JWTTokenService) ParseRefreshToken(tokenStr string) (domUser.UserID, string, error) {
	var claims refreshClaims
	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return s.refreshSecret, nil
	})
	if err != nil {
		return 0, "", err
	}
	return claims.UserID, claims.TokenID, nil
}
