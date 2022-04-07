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

func getMigrationFile(name string) MigrationFile {
	for _, mfile := range migrationFiles {
		if mfile.FileName == name {
			return mfile
		}
	}
	return MigrationFile{}
}

func (mf *MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mf.FileName {
			return false
		}
	}
	return true
}
