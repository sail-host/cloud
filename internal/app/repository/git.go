package repository

import (
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
)

type GitRepo struct {
}

type IGitRepo interface {
	GetGitByID(id uint) (*model.Git, error)
	GetListGit() ([]model.Git, error)
	CreateGit(git *model.Git) error
	UpdateGit(git *model.Git) error
	DeleteGit(id uint) error
}

func NewIGitRepo() IGitRepo {
	return &GitRepo{}
}

func (g *GitRepo) GetGitByID(id uint) (*model.Git, error) {
	var git model.Git
	db := global.DB
	err := db.Where("id = ?", id).First(&git).Error
	return &git, err
}

func (g *GitRepo) GetListGit() ([]model.Git, error) {
	var git []model.Git
	db := global.DB
	err := db.Find(&git).Error
	return git, err
}

func (g *GitRepo) CreateGit(git *model.Git) error {
	db := global.DB
	err := db.Create(git).Error
	return err
}

func (g *GitRepo) UpdateGit(git *model.Git) error {
	db := global.DB
	err := db.Model(git).Updates(git).Error
	return err
}

func (g *GitRepo) DeleteGit(id uint) error {
	db := global.DB
	err := db.Where("id = ?", id).Delete(&model.Git{}).Error
	return err
}
