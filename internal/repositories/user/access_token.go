package userrepository

import (
	"context"
	"errors"
	"time"

	"bookingrooms/internal/models"
)

func (r repository) RevokedAccessToken(idToken string) error {
	count, err := r.GetAccessTokenByIdToken(idToken)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("no data")
	}

	return r.db.
		Model(&models.AccessTokens{}).
		Where("id_token = ?", idToken).
		Update("revoked", "1").
		Error
}

func (r repository) GetAccessTokenByIdToken(idToken string) (int, error) {
	var count int
	return count, r.db.
		Model(&models.AccessTokens{}).
		Select("COUNT(*)").
		Where("id_token = ?", idToken).
		First(&count).Error
}

func (r repository) GetRevokedToken(id string) (string, error) {
	var revoked string
	return revoked, r.db.
		Model(&models.AccessTokens{}).
		Select("revoked").
		Where("id_token = ?", id).
		First(&revoked).Error
}

func (r repository) SetRevokedTokenToRedis(idToken, revoked string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	return r.rdb.Set(ctx, idToken, revoked, 24*time.Hour).Err()
}

func (r repository) GetRevokedTokenFromRedis(idToken string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	return r.rdb.Get(ctx, idToken).Result()
}
