package main

import (
	"fmt"
	"gohub/app/cmd"
	"gohub/app/cmd/make"
	"gohub/bootstrap"
	bstConfig "gohub/config"
	"gohub/pkg/config"
	"gohub/pkg/console"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	bstConfig.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple forum project",
		Long: `Default will run "serve" command, you can use "-h" flag to see
		 all subcommands`,
		PersistentPreRun: func(command *cobra.Command, args []string) {
			config.InitConfig(cmd.Env)
			bootstrap.SetupLogger()
			bootstrap.SetupDB()
			bootstrap.SetupRedis()
			bootstrap.SetupCache()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,
	)

	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
