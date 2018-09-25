package cmd

import (
	"io"
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"
)

func newDeletCmd(client *grafana.Client, out io.Writer) *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete a resource by name or id",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
			return nil
		},
	}

	deleteCmd.AddCommand(newDatasourceDeleteCommand(client, out))
	deleteCmd.AddCommand(newDashboardDeleteCommand(client, out))
	return deleteCmd
}
