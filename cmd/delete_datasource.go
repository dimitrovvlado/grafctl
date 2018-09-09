package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/dimitrovvlado/grafctl/grafana"

	"github.com/spf13/cobra"
)

type datasourceDeleteCmd struct {
	client       *grafana.Client
	out          io.Writer
	datasourceID string
}

func newDatasourceDeleteCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	i := &datasourceDeleteCmd{
		client: client,
		out:    out,
	}
	deleteDatasourcesCmd := &cobra.Command{
		Use:     "datasource",
		Aliases: []string{"ds"},
		Short:   "Delete datasource by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				log.Println("Command 'delete' requires an ID")
				cmd.Help()
				return nil
			}
			ensureClient(i.client)
			i.datasourceID = args[0]
			return i.run()
		},
	}

	return deleteDatasourcesCmd
}

// run deletes a ds
func (i *datasourceDeleteCmd) run() error {
	ds, err := i.client.GetDatasource(i.datasourceID)
	if err != nil {
		log.Fatalln(err)
	}
	err = i.client.DeleteDatasource(ds)
	if err != nil {
		return err
	}
	fmt.Fprintln(i.out, fmt.Sprintf("Datasource '%s' deleted.", ds.Name))
	return nil
}
