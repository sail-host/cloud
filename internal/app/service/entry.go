package service

import (
	repo "github.com/sail-host/cloud/internal/app/repository"
)

var (
	userRepo = repo.NewIUserRepo()
	authRepo = repo.NewIAuthRepo()
	gitRepo  = repo.NewIGitRepo()
)