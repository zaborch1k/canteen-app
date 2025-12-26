package jwtadapter

import (
	domAuth "canteen-app/internal/domain/auth"
	domUser "canteen-app/internal/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenService struct {
	secret []byte
	issuer string
}

func NewJWTTokenService(secret []byte, issuer string) *JWTTokenService {
	return &JWTTokenService{secret: secret, issuer: issuer}
}

type jwtClaims struct {
	UserID domUser.UserID `json:"sub"`
	Role   string         `json:"role"`
	jwt.RegisteredClaims
}

func (s *JWTTokenService) GenerateAccesToken(c domAuth.Claims) (string, error) {
	claims := jwtClaims{
		UserID: c.UserID,
		Role:   c.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(c.ExpireAt),
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(s.secret)
}
