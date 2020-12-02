package platform

import (
	"github.com/spf13/cobra"
)

// ResponsePayload for receiving the request

// StartCmd subcommand for Airy Core
var StartCmd = &cobra.Command{
	Use:              "start",
	TraverseChildren: true,
	Short:            "Reset apps",
	Long: `Reset an airy app, with the following order:
- scale down the statefulset or deployment
- delete the pvc (unless specified to keep it)
- wait for the consumer group to be empty and then reset the offset on some or all normal topics, to either earliest or latest
- delte the existing schema (unless specified to keep it)`,
	Run: start,
}

func start(cmd *cobra.Command, args []string) {
	// Initialize the api request

}

func init() {
	var appname string
	StartCmd.Flags().StringVarP(&appname, "app-name", "a", "", "The app which we will reset.")
}
