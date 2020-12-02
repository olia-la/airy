package core

import (
	auth "cli/cmd/core/auth"
	core "cli/cmd/core/start"

	"github.com/spf13/cobra"
)

// CoreCmd root command for managing apps with the cli
var CoreCmd = &cobra.Command{
	Use:              "core",
	TraverseChildren: true,
	Short:            "Manage Airy Core Platform",
	Long:             `Run complex operations for the apps we deploy with k8s & terraform`,
}

func init() {

	CoreCmd.AddCommand(core.StartCmd)
	CoreCmd.AddCommand(auth.AuthCmd)
}
