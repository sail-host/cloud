package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/init/migration/migrations"
)

func Init() {
	m := gormigrate.New(global.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrations.CreateUserTable,
		migrations.CreateAuthTokenTable,
		migrations.CreateGitTable,
		migrations.CreateDomainTable,
		migrations.CreateProjectTable,
		migrations.CreateEnvironmentVariableTable,
		migrations.CreateDeploymentTable,
		migrations.CreateLogTable,
		migrations.CreateProjectDomainTable,
	})

	if err := m.Migrate(); err != nil {
		global.LOG.Error(err)
		panic(err)
	}
	global.LOG.Info("Migration run successfully")
}
