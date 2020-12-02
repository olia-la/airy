package platform

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ResponsePayload for receiving the request

// StartCmd subcommand for Airy Core
var StartCmd = &cobra.Command{
	Use:              "start",
	TraverseChildren: true,
	Short:            "Start Airy Core Platform",
	Long: `This will install the Airy Core Platform in the current directory unless you choose a different one.
	It will also try to install Vagrant and VirtualBox.`,
	Run: start,
}

func start(cmd *cobra.Command, args []string) {
	// Initialize the api request

	fmt.Println("Starting bootstrap")

}

func init() {
	var appname string
	StartCmd.Flags().StringVarP(&appname, "app-name", "a", "", "The app which we will reset.")
}
