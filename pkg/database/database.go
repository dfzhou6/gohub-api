package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gohub/pkg/config"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger gormLogger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() string {
	return DB.Migrator().CurrentDatabase()
}

func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMySQLTables()
	default:
		panic(errors.New("database connection not supported"))
	}
	return err
}

func deleteMySQLTables() error {
	dbName := CurrentDatabase()
	tables := []string{}
	err := DB.Table("information_schema.tables").
		Where("table_schema", dbName).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	DB.Exec("SET foreign_key_checks = 0;")

	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}

func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(obj)
	return stmt.Schema.Table
}
