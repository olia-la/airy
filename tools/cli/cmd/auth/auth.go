package auth

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ResponsePayload for receiving the request

// AuthCmd subcommand for Airy Core
var AuthCmd = &cobra.Command{
	Use:              "auth",
	TraverseChildren: true,
	Short:            "Create a default user and return a JWT token",
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
