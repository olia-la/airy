package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type loginRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponsePayload struct {
	ID        string
	FirstName string
	LastName  string
	Token     string
}

// AuthCmd subcommand for Airy Core
var AuthCmd = &cobra.Command{
	Use:              "auth",
	TraverseChildren: true,
	Short:            "Create a default user and return a JWT token",
	Long:             ``,
	Run:              auth,
}

func sendRequest(requestDataJSON []byte, url string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestDataJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Response Status code is %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)

}

func auth(cmd *cobra.Command, args []string) {
	url := "http://api.airy/users.login"

	requestPayload := loginRequestPayload{Email: "grace@example.com", Password: "the_answer_is_42"}
	requestDataJSON, err := json.Marshal(requestPayload)
	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := sendRequest(requestDataJSON, url)
	if err != nil {
		log.Fatal(err)
	}

	var loginResponsePayload loginResponsePayload
	jsonErr := json.Unmarshal(responseBody, &loginResponsePayload)
	if jsonErr != nil {
		log.Fatal("Error unmarshaling response")
	}
	fmt.Println(loginResponsePayload.Token)
}

func init() {
	var appname string
	AuthCmd.Flags().StringVarP(&appname, "app-name", "a", "", "The app which we will reset.")
}
