package main

import (
	"os"
	"os/exec"
	"testing"

	airytests "cli/pkg/tests"
	"reflect"
)

const binaryName = "./airy"

func TestCli(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		golden  string
		wantErr bool
	}{
		{"no args", []string{}, "cli.no-args.golden", false},
		{"kafka status", []string{"kafka", "status", "--airyUrl", "http://localhost:3001"}, "cli.kafka.status.golden", false},
		{"kafka consumergroup users", []string{"kafka", "consumergroup", "--airyUrl", "http://localhost:3001"}, "cli.kafka.consumergroup.golden", false},
		{"app reset no args", []string{"app", "reset"}, "cli.app.reset.no-args.golden", true},
		{"app reset no app-name", []string{"app", "reset", "--app-name", "ops-test"}, "cli.app.reset.no-args.golden", true},
		{"app reset no app-id", []string{"app", "reset", "--app-id", "ops-test"}, "cli.app.reset.no-args.golden", true},
	}

	go func() {
		airytests.MockServer()
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(testing *testing.T) {
			cmd := exec.Command(binaryName, tt.args...)
			os.Setenv("AWS_PROFILE", "airy-mock")
			output, err := cmd.CombinedOutput()

			if (err != nil) != tt.wantErr {
				t.Fatalf("Test expected to fail: %t. Did the test pass: %t. Error message: %v\n", tt.wantErr, err == nil, err)
			}

			actual := string(output)
			golden := airytests.NewGoldenFile(t, tt.golden)
			expected := golden.Load()

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("diff: %v", airytests.Diff(actual, expected))
			}

		})

	}
}
