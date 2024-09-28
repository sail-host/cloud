package repository

import (
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
)

type UserRepo struct {
}

type IUserRepo interface {
	GetUserByEmail(email string) (*model.User, error)
	CheckUsersExist() (bool, error)
	CreateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
}

func NewIUserRepo() IUserRepo {
	return &UserRepo{}
}

func (u *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	db := global.DB
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepo) CheckUsersExist() (bool, error) {
	var count int64
	db := global.DB
	err := db.Model(&model.User{}).Count(&count).Error
	return count > 0, err
}

func (u *UserRepo) CreateUser(user *model.User) error {
	db := global.DB
	err := db.Create(user).Error
	return err
}

func (u *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	db := global.DB
	err := db.Where("id = ?", id).First(&user).Error
	return &user, err
}
