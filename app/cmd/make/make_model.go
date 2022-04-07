package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1),
}

func runMakeModel(command *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	dir := fmt.Sprintf("app/models/%s/", model.PackageName)
	os.MkdirAll(dir, os.ModePerm)

	createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
	createFileFromStub(dir+model.PackageName+"_util.go", "model/model_util", model)
	createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)
}
