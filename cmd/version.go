package cmd

import (
	"fmt"

	"github.com/dimitrovvlado/grafctl/version"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of grafctl",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Grafana command tool %s\n", version.VERSION)
		},
	}

	return versionCmd
}
