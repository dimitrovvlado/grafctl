package cmd

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

type dashboardListCmd struct {
	// ListDashboards
	client *grafana.Client
	out    io.Writer
	output string
}

func newDashboardListCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	get := &dashboardListCmd{
		client: client,
		out:    out,
	}
	getDashboardsCmd := &cobra.Command{
		Use:     "dashboards",
		Aliases: []string{"db", "dbs"},
		Short:   "Display one or many dashboards",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(get.client)
			return get.run()
		},
	}
	f := getDashboardsCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
	return getDashboardsCmd
}

func (i *dashboardListCmd) run() error {
	db, err := i.client.ListDashboards()
	if err != nil {
		log.Fatalln(err)
	}

	//TODO extract as flag
	var colWidth uint = 60
	formatter := func() string {
		if db == nil || len(db) == 0 {
			return fmt.Sprintf("No dashboards found.")
		}
		table := uitable.New()
		table.MaxColWidth = colWidth
		table.AddRow("ID", "TITLE", "URI", "TAGS")
		for _, lr := range db {
			table.AddRow(lr.ID, lr.Title, lr.URI, strings.Join(lr.Tags, ", "))
		}
		return fmt.Sprintf("%s%s", table.String(), "\n")
	}

	result, err := formatResult(i.output, db, formatter)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(i.out, result)

	return nil
}
