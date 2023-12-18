package migrations

import (
	"gofr.dev/cmd/gofr/migration/dbMigration"
)

func All() map[string]dbmigration.Migrator {
	return map[string]dbmigration.Migrator{

		"20231210055442": K20231210055442{},
	}
}
