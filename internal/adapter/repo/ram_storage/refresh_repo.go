package ram_storage

import (
	"time"

	domUser "canteen-app/internal/domain/user"
	"canteen-app/internal/usecase"
)

type refreshRecord struct {
	UserId    domUser.UserID
	ExpiresAt time.Time
}

type RefreshRepo struct {
	data map[string]refreshRecord
}

var _ usecase.RefreshTokenRepository = (*RefreshRepo)(nil)

func NewRefreshRepo() *RefreshRepo {
	return &RefreshRepo{
		data: make(map[string]refreshRecord),
	}
}

func (r *RefreshRepo) Save(tokenID string, userID domUser.UserID, exp time.Time) {
	r.data[tokenID] = refreshRecord{UserId: userID, ExpiresAt: exp}
}

func (r *RefreshRepo) Delete(tokenID string) {
	delete(r.data, tokenID)
}

func (r *RefreshRepo) IsValid(tokenID string, userID domUser.UserID) bool {
	rec, ok := r.data[tokenID]
	if !ok {
		return false
	}
	if rec.UserId != userID {
		return false
	}
	return true
}
