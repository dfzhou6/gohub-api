package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

type migrationFunc func(gorm.Migrator, *sql.DB)

type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

var migrationFiles []MigrationFile

func Add(name string, up, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		Up:       up,
		Down:     down,
		FileName: name,
	})
}
