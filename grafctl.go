package main

import (
	"fmt"
	"os"

	"github.com/dimitrovvlado/grafctl/cmd"
	"github.com/dimitrovvlado/grafctl/grafana"
)

var (
	settings EnvSettings
	client   *grafana.Client
)

func main() {
	settings.Init()

	if len(os.Args) > 1 {
		// Create the Grafana client.
		client = grafana.New(settings.GrafanaHost, settings.GrafanaUsername, settings.GrafanaPassword)
	}

	rootCmd := cmd.NewRootCmd(client)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// CheckIfErrorAndExit should be used to naively panics if an error is not nil.
func CheckIfErrorAndExit(err error) {
	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
