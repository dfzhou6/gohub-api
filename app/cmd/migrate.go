package cmd

import (
	"gohub/database/migrations"
	"gohub/pkg/migrate"

	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

var CmdMigrateDown = &cobra.Command{
	Use:     "down",
	Short:   "Reverse the up command",
	Aliases: []string{"rollback"},
	Run:     runDown,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateDown,
	)
}

func migrator() *migrate.Migrator {
	migrations.Initialize()
	return migrate.NewMigrator()
}

func runUp(command *cobra.Command, args []string) {
	migrator().Up()
}

func runDown(command *cobra.Command, args []string) {
	migrator().Rollback()
}
