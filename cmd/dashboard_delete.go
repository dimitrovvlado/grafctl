package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"
)

type dashboardDeleteCmd struct {
	client *grafana.Client
	out    io.Writer
	uuid   string
	output string
}

func newDashboardDeleteCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	del := &dashboardDeleteCmd{
		client: client,
		out:    out,
	}
	deleteDashboardCmd := &cobra.Command{
		Use:     "dashboard",
		Aliases: []string{"db", "dbs", "dashboards"},
		Short:   "Delete dashboard by UID",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(del.client)
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
			del.uuid = args[0]
			return del.run()
		},
	}
	return deleteDashboardCmd
}

func (i *dashboardDeleteCmd) run() error {
	title, err := i.client.DeleteDashboard(i.uuid)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(i.out, fmt.Sprintf("Dashboard \"%s\" deleted", title))
	return nil
}
