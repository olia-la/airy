package cmd

import "github.com/spf13/cobra"

// StatusCmd cli kafka version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return kafka broker version",
	Long: `We inspect version in both the kafka brokers and Zookeeper.

Controller: Is the broker the current controller or not (there must be only one)
KafkaStatus: What is the version of the broker, when we ask him directly
ZookeeperStatus: What does Zookeeper say about the health of the broker
`,
	Run: version,
}

func version(cmd *cobra.Command, args []string) {

}

func init() {
	CoreCmd.AddCommand(versionCmd)
}
