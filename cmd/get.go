package cmd

import (
	"io"
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"
)

func newGetCmd(client *grafana.Client, out io.Writer) *cobra.Command {
	getCmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{""},
		Short:   "Display one or many resources",
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
			return nil
		},
	}

	getCmd.AddCommand(newOrgListCommand(client, out))
	getCmd.AddCommand(newUsersListCommand(client, out))
	getCmd.AddCommand(newDatasourceListCommand(client, out))

	return getCmd
}
