package domain

type UserID int64

type User struct {
	ID           UserID
	Login        string
	PasswordHash string
	Role         string
}
