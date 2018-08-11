package cmd

import (
	"fmt"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

type datasourceCmd struct {
	client *grafana.Client
	output string
}

func newDatasourceCommand(client *grafana.Client) *cobra.Command {
	get := &datasourceCmd{
		client: client,
	}
	getDatasourcesCmd := &cobra.Command{
		Use:     "datasources",
		Aliases: []string{"ds"},
		Short:   "Display one or many datasources",
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			get.output = cmd.Flag("output").Value.String()
			return get.run()
		},
	}
	return getDatasourcesCmd
}

// run creates a merge request
func (i *datasourceCmd) run() error {
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
	fmt.Printf(result)

	return nil
}
