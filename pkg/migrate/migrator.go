package migrate

import (
	"gohub/pkg/console"
	"gohub/pkg/database"
	"gohub/pkg/file"
	"io/ioutil"

	"gorm.io/gorm"
)

type Migrator struct {
	Folder    string
	DB        *gorm.DB
	GMigrator gorm.Migrator
}

type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

func NewMigrator() *Migrator {
	m := &Migrator{
		Folder:    "database/migrations/",
		DB:        database.DB,
		GMigrator: database.DB.Migrator(),
	}

	m.createMigrationsTable()
	return m
}

func (m *Migrator) createMigrationsTable() {
	migration := Migration{}
	if !m.GMigrator.HasTable(&migration) {
		m.GMigrator.CreateTable(&migration)
	}
}

func (m *Migrator) readAllMigrationFiles() []MigrationFile {
	files, err := ioutil.ReadDir(m.Folder)
	console.ExitIf(err)

	var migrationFiles []MigrationFile
	for _, f := range files {
		fileName := file.FileNameWithoutExtension(f.Name())
		mfile := getMigrationFile(fileName)
		if len(mfile.FileName) > 0 {
			migrationFiles = append(migrationFiles, mfile)
		}
	}

	return migrationFiles
}

func (m *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	if mfile.Up != nil {
		console.Warning("migrating " + mfile.FileName)
		mfile.Up(m.GMigrator, database.SQLDB)
		console.Success("migrated " + mfile.FileName)
	}

	err := m.DB.Create(&Migration{
		Migration: mfile.FileName,
		Batch:     batch,
	}).Error

	console.ExitIf(err)
}

func (m *Migrator) getBatch() int {
	batch := 1
	lastMigration := Migration{}
	m.DB.Order("id desc").First(&lastMigration)
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

func (m *Migrator) Up() {
	migrationFiles := m.readAllMigrationFiles()
	batch := m.getBatch()

	migrations := []Migration{}
	m.DB.Find(&migrations)

	runed := false
	for _, mfile := range migrationFiles {
		if mfile.isNotMigrated(migrations) {
			m.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}
}

func (m *Migrator) rollbackMigrations(migrations []Migration) bool {
	runed := false
	for _, migration := range migrations {
		console.Warning("rollback " + migration.Migration)
		mfile := getMigrationFile(migration.Migration)
		if mfile.Down != nil {
			mfile.Down(m.GMigrator, database.SQLDB)
		}

		runed = true

		m.DB.Delete(&migration)

		console.Success("finish " + mfile.FileName)
	}
	return runed
}

func (m *Migrator) Rollback() {
	lastMigration := Migration{}
	m.DB.Order("id desc").First(&lastMigration)
	migrations := []Migration{}
	m.DB.Where("batch", lastMigration.Batch).Order("id desc").Find(&migrations)
	if !m.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}
