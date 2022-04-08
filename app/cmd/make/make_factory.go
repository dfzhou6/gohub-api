package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, exmaple: make factory user",
	Args:  cobra.ExactArgs(1),
	Run:   runMakeFactory,
}

func runMakeFactory(command *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	filePath := fmt.Sprintf("database/factories/%s_factory.go", model.PackageName)
	createFileFromStub(filePath, "factory", model)
}
