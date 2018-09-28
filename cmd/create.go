package cmd

import (
	"io"
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"
)

func newCreateCmd(client *grafana.Client, out io.Writer) *cobra.Command {
	createCmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{""},
		Short:   "Create a resource",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
			return nil
		},
	}

	createCmd.AddCommand(newDatasourceCreateCommand(client, out))
	createCmd.AddCommand(newDashboardCreateCommand(client, out))
	createCmd.AddCommand(newTeamCreateCommand(client, out))

	return createCmd
}
