package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.AddCommand(versionCmd)
}

func newVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of grafctl",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Grafana command tool v0.1")
		},
	}

	return versionCmd
}
