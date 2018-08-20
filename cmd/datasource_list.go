package cmd

import (
	"fmt"
	"io"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

type datasourceListCmd struct {
	client *grafana.Client
	out    io.Writer
	output string
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
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return get.run()
		},
	}
	f := getDatasourcesCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
	return getDatasourcesCmd
}

// run creates a merge request
func (i *datasourceListCmd) run() error {
	ds, err := i.client.ListDatasources()
	if err != nil {
		logrus.Fatal(err)
	}

	//TODO extract as flag
	var colWidth uint = 60
	formatter := func() string {
		table := uitable.New()
		table.MaxColWidth = colWidth
		table.AddRow("ID", "NAME", "TYPE", "ACCESS", "URL")
		for _, lr := range ds {
			table.AddRow(lr.ID, lr.Name, lr.Type, lr.Access, lr.URL)
		}
		return fmt.Sprintf("%s%s", table.String(), "\n")
	}

	result, err := formatResult(i.output, ds, formatter)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Fprintln(i.out, result)

	return nil
}
