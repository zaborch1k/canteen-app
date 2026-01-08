package password

import (
	"canteen-app/internal/usecase"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

var _ usecase.PasswordHasher = BcryptHasher{}

func (h BcryptHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (h BcryptHasher) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
