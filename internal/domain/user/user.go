package user

import "golang.org/x/crypto/bcrypt"

type UserID int64

type User struct {
	ID           UserID
	Login        string
	PasswordHash string
	Name         string
	Surname      string
	Role         string
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
