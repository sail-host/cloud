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
