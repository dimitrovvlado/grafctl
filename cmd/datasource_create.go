package cmd

import (
	"encoding/json"
	"io/ioutil"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

type datasourceCreateCmd struct {
	client *grafana.Client
	output string
	file   string
}

func newDatasourceCreateCommand(client *grafana.Client) *cobra.Command {
	i := &datasourceCreateCmd{
		client: client,
	}
	createDatasourcesCmd := &cobra.Command{
		Use:     "datasource",
		Aliases: []string{"ds"},
		Short:   "Create datasource",
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			i.output = cmd.Flag("output").Value.String()
			if i.file == "" {
				logrus.Warn("Command needs either a file reference or a set ot values.")
				cmd.Help()
				return nil
			}

			return i.run()
		},
	}

	createDatasourcesCmd.PersistentFlags().StringVarP(&i.file, "from-file", "f", "", "A json file with the datasource")

	return createDatasourcesCmd
}

// run creates a merge request
func (i *datasourceCreateCmd) run() error {

	var datasource grafana.Datasource
	byteValue, _ := ioutil.ReadFile(i.file)

	err := json.Unmarshal(byteValue, &datasource)
	if err != nil {
		logrus.Fatal(err)
	}

	ds, err := i.client.CreateDatasource(datasource)
	if err != nil {
		logrus.Fatal(err)
	}
	//TODO
	logrus.Info(ds.ID)

	return nil
}
