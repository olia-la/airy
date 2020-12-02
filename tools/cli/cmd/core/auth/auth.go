package platform

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ResponsePayload for receiving the request

// StartCmd subcommand for Airy Core
var AuthCmd = &cobra.Command{
	Use:              "auth",
	TraverseChildren: true,
	Short:            "Creates a default user and returns a JWT token.",
	Long:             ``,
	Run:              auth,
}

func auth(cmd *cobra.Command, args []string) {
	// Initialize the api request

	fmt.Println("Valid Token")

}

func init() {
	var appname string
	AuthCmd.Flags().StringVarP(&appname, "app-name", "a", "", "The app which we will reset.")
}
