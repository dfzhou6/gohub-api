package migrate

import (
	"gohub/pkg/database"

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
