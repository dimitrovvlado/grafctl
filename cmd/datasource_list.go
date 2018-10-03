package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"
)

type datasourceListCmd struct {
	client *grafana.Client
	out    io.Writer
	output string
	ID     string
}

func newDatasourceListCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	get := &datasourceListCmd{
		client: client,
		out:    out,
	}
	getDatasourcesCmd := &cobra.Command{
		Use:     "datasources",
		Aliases: []string{"ds"},
		Short:   "Display one or many datasources",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(get.client)
			if len(args) > 0 {
				get.ID = args[0]
			}
			return get.run()
		},
	}
	f := getDatasourcesCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
	return getDatasourcesCmd
}

func (i *datasourceListCmd) run() error {
	var colWidth uint = 60
	var obj interface{}
	var formatter func() string

	if i.ID != "" {
		data, err := i.client.GetDatasource(i.ID)
		if err != nil {
			log.Fatalln(err)
		}
		obj = data
		formatter = func() string {
			if (grafana.Datasource{}) == obj {
				return fmt.Sprintf("Datasource not found")
			}
			table := uitable.New()
			table.MaxColWidth = colWidth
			table.AddRow("ID", "NAME", "TYPE", "ACCESS", "URL")
			table.AddRow(data.ID, data.Name, data.Type, data.Access, data.URL)
			return fmt.Sprintf("%s\n", table.String())
		}
	} else {
		data, err := i.client.ListDatasources()
		if err != nil {
			log.Fatalln(err)
		}
		obj = data
		formatter = func() string {
			if obj == nil || len(data) == 0 {
				return fmt.Sprintf("No datasources found")
			}
			table := uitable.New()
			table.MaxColWidth = colWidth
			table.AddRow("ID", "NAME", "TYPE", "ACCESS", "URL")
			for _, ds := range data {
				table.AddRow(ds.ID, ds.Name, ds.Type, ds.Access, ds.URL)
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
