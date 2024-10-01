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

var CreateProjectTable = &gormigrate.Migration{
	ID: "20241024-add-table-project",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.Project{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&model.Project{})
	},
}

var CreateEnvironmentVariableTable = &gormigrate.Migration{
	ID: "20241024-add-table-environment-variable",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.EnvironmentVariable{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&model.EnvironmentVariable{})
	},
}

var CreateDeploymentTable = &gormigrate.Migration{
	ID: "20241024-add-table-deployment",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.Deployment{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&model.Deployment{})
	},
}

var CreateLogTable = &gormigrate.Migration{
	ID: "20241024-add-table-log",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.Log{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&model.Log{})
	},
}

var CreateProjectDomainTable = &gormigrate.Migration{
	ID: "20241024-add-table-project-domain",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&model.ProjectDomain{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&model.ProjectDomain{})
	},
}
