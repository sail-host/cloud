package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/sail-host/cloud/internal/app/model"
	"gorm.io/gorm"
)

var CreateUserTable = &gormigrate.Migration{
	ID: "20241024-add-table-user",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&model.User{})
	},
}

var CreateAuthTokenTable = &gormigrate.Migration{
	ID: "20241024-add-table-auth-token",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.AuthToken{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&model.AuthToken{})
	},
}

var CreateGitTable = &gormigrate.Migration{
	ID: "20241024-add-table-git",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.Git{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&model.Git{})
	},
}

var CreateDomainTable = &gormigrate.Migration{
	ID: "20241024-add-table-domain",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.Domain{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&model.Domain{})
	},
}
