package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

// StatusCmd cli kafka version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return current version",
	Long: `We inspect version in both the kafka brokers and Zookeeper.

Controller: Is the broker the current controller or not (there must be only one)
KafkaStatus: What is the version of the broker, when we ask him directly
ZookeeperStatus: What does Zookeeper say about the health of the broker
`,
	Run: version,
}

func version(cmd *cobra.Command, args []string) {
	airyHomePath := "/Users/pascal/core/"
	content, err := ioutil.ReadFile(airyHomePath + "VERSION")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
