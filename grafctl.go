package main

import (
	"log"
	"os"

	"github.com/dimitrovvlado/grafctl/cmd"
	"github.com/dimitrovvlado/grafctl/environment"
	"github.com/dimitrovvlado/grafctl/grafana"
)

var (
	settings environment.EnvSettings
	client   *grafana.Client
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	err := settings.Init()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 && settings.Initialized {
		// Create the Grafana client.
		client = grafana.New(settings.GrafanaHost, settings.GrafanaUsername, settings.GrafanaPassword)
	}

	rootCmd := cmd.NewRootCmd(client)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
