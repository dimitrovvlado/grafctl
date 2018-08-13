package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

type datasourceCreateCmd struct {
	client *grafana.Client
	output string
	files  *[]string
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
			//TODO maybe change the command to create a ds by a given set ot flags
			if i.files == nil && len(*i.files) == 0 {
				logrus.Warn("Command needs either a file reference or a set ot values.")
				cmd.Help()
				return nil
			}

			return i.run()
		},
	}

	i.files = createDatasourcesCmd.PersistentFlags().StringSliceP("filename", "f", []string{}, "Filename(s) or direcory to use to create the datasource")
	return createDatasourcesCmd
}

// run creates a merge request
func (i *datasourceCreateCmd) run() error {
	for _, file := range *i.files {
		importDatasource(file, i.client)
	}
	return nil
}

func importDatasource(filename string, client *grafana.Client) {
	info, err := os.Stat(filename)
	if err != nil {
		logrus.Warn(err)
		return
	}
	if info.IsDir() {
		files, err := filePathWalkDir(filename)
		if err != nil {
			logrus.Warn(err)
		} else {
			for _, fi := range files {
				importDatasource(fi, client)
			}
		}
	} else {
		var datasource grafana.Datasource
		byteValue, _ := ioutil.ReadFile(filename)

		err := json.Unmarshal(byteValue, &datasource)
		if err != nil {
			logrus.Warn(err)
		}

		ds, err := client.CreateDatasource(datasource)
		if err != nil {
			logrus.Warn(err)
		} else {
			logrus.Info("Datasource \"", ds.Name, "\" created")
		}
	}
}

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
