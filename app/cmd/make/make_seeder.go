package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "Create seeder file, example:  make seeder user",
	Args:  cobra.ExactArgs(1),
	Run:   runMakeSeeder,
}

func runMakeSeeder(command *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	filePath := fmt.Sprintf("database/seeders/%s_seeder.go", model.TableName)
	createFileFromStub(filePath, "seeder", model)
}
