package main

import (
	"cli/cmd"
)

var Version string

func main() {

	cmd.VersionString = Version
	cmd.Execute()
}
