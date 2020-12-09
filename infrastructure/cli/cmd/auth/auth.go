package auth

import (
	"fmt"

	"github.com/airyhq/airy/lib/go/apiclient"
	"github.com/airyhq/airy/lib/go/apiclient/payloads"

	"github.com/spf13/cobra"
)

// AuthCmd subcommand for Airy Core
var AuthCmd = &cobra.Command{
	Use:              "auth",
	TraverseChildren: true,
	Short:            "Create a default user and return a JWT token",
	Long:             ``,
	Run:              auth,
}

func auth(cmd *cobra.Command, args []string) {
	c := apiclient.NewClient()
	loginRequestPayload := payloads.LoginRequestPayload{Email: "grace@example.com", Password: "the_answer_is_42"}

	res, err := c.Login(loginRequestPayload)
	fmt.Println(res.Token)
}

func init() {
}
