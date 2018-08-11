package cmd

import (
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"
)

func newGetCmd(client *grafana.Client) *cobra.Command {
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

	getCmd.AddCommand(newOrgCommand(client))
	getCmd.AddCommand(newUsersCommand(client))
	getCmd.AddCommand(newDatasourceCommand(client))

	return getCmd
}
