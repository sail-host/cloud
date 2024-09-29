package repository

import (
	"time"

	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
)

type AuthRepo struct {
}

type IAuthRepo interface {
	CreateAuthToken(authToken *model.AuthToken) error
	GetAuthTokenByToken(token string) (*model.AuthToken, error)
	DeleteAuthTokenByToken(token string) error
	DeleteAuthTokenByUserID(userID uint) error
	DeleteExpiredAuthTokens() error
}

func NewIAuthRepo() IAuthRepo {
	return &AuthRepo{}
}

func (a *AuthRepo) CreateAuthToken(authToken *model.AuthToken) error {
	db := global.DB
	return db.Create(authToken).Error
}

func (a *AuthRepo) GetAuthTokenByToken(token string) (*model.AuthToken, error) {
	var authToken model.AuthToken
	db := global.DB
	err := db.Where("token = ?", token).First(&authToken).Error
	return &authToken, err
}

func (a *AuthRepo) DeleteAuthTokenByToken(token string) error {
	db := global.DB
	return db.Where("token = ?", token).Delete(&model.AuthToken{}).Error
}

func (a *AuthRepo) DeleteAuthTokenByUserID(userID uint) error {
	db := global.DB
	return db.Where("user_id = ?", userID).Delete(&model.AuthToken{}).Error
}

func (a *AuthRepo) DeleteExpiredAuthTokens() error {
	db := global.DB
	return db.Where("expired_at < ?", time.Now()).Delete(&model.AuthToken{}).Error
}
