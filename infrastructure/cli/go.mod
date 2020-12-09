module cli

go 1.12

require (
	apiclient
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.1
)

replace apiclient => ../../lib/go/apiclient
