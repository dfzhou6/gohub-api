package make

import (
	"fmt"
	"gohub/pkg/console"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller, exmaple: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(command *cobra.Command, args []string) {
	arr := strings.Split(args[0], "/")
	if len(arr) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	apiVersion, name := arr[0], arr[1]
	model := makeModelFromString(name)

	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)

	createFileFromStub(filePath, "apicontroller", model)
}
