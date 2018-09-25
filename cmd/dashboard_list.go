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
	uuid   string
	output string
}

func newDashboardListCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	get := &dashboardListCmd{
		client: client,
		out:    out,
	}
	getDashboardsCmd := &cobra.Command{
		Use:     "dashboards",
		Aliases: []string{"db", "dbs", "dashboard"},
		Short:   "Display one or many dashboards",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(get.client)
			if len(args) > 0 {
				get.uuid = args[0]
			}
			return get.run()
		},
	}
	f := getDashboardsCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
	return getDashboardsCmd
}

func (i *dashboardListCmd) run() error {
	//TODO extract as flag
	var colWidth uint = 60
	var obj interface{}
	var formatter func() string
	if i.uuid != "" {
		ex, err := i.client.GetDashboard(i.uuid)
		if err != nil {
			log.Fatalln(err)
		}
		formatter = func() string {
			if (grafana.DashboardExport{}) == obj {
				return fmt.Sprintf("Dashboard not found")
			}
			table := uitable.New()
			table.MaxColWidth = colWidth
			table.AddRow("Slug", "Version", "Folder", "URL")
			table.AddRow(ex.Meta.Slug, ex.Meta.Version, ex.Meta.FolderTitle, ex.Meta.URL)
			return fmt.Sprintf("%s\n", table.String())
		}
		obj = ex.Dashboard
	} else {
		obj, err := i.client.ListDashboards()
		if err != nil {
			log.Fatalln(err)
		}
		formatter = func() string {
			if obj == nil || len(obj) == 0 {
				return fmt.Sprintf("No dashboards found.")
			}
			table := uitable.New()
			table.MaxColWidth = colWidth
			table.AddRow("ID", "UID", "TITLE", "URI", "TAGS")
			for _, lr := range obj {
				table.AddRow(lr.ID, lr.UID, lr.Title, lr.URI, strings.Join(lr.Tags, ", "))
			}
			return fmt.Sprintf("%s\n", table.String())
		}
	}

	result, err := formatResult(i.output, obj, formatter)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(i.out, result)

	return nil
}
