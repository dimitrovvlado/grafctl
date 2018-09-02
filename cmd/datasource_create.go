package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dimitrovvlado/grafctl/grafana"

	"github.com/spf13/cobra"
)

type datasourceCreateCmd struct {
	client *grafana.Client
	out    io.Writer
	files  *[]string
}

func newDatasourceCreateCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	i := &datasourceCreateCmd{
		client: client,
		out:    out,
	}
	createDatasourcesCmd := &cobra.Command{
		Use:     "datasource",
		Aliases: []string{"ds"},
		Short:   "Create datasource",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(i.client)
			//TODO maybe change the command to create a ds by a given set ot flags
			return i.run()
		},
	}

	i.files = createDatasourcesCmd.PersistentFlags().StringSliceP("filename", "f", []string{}, "Filename(s) or direcory to use to create the datasource")
	createDatasourcesCmd.MarkPersistentFlagRequired("filename")
	return createDatasourcesCmd
}

// run creates a datasource
func (i *datasourceCreateCmd) run() error {
	for _, file := range *i.files {
		importDatasource(file, i)
	}
	return nil
}

func importDatasource(filename string, cmd *datasourceCreateCmd) {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintln(cmd.out, err)
		return
	}
	if info.IsDir() {
		files, err := filePathWalkDir(filename)
		if err != nil {
			fmt.Fprintln(cmd.out, err)
		} else {
			for _, fi := range files {
				importDatasource(fi, cmd)
			}
		}
	} else {
		var datasource grafana.Datasource
		byteValue, _ := ioutil.ReadFile(filename)

		err := json.Unmarshal(byteValue, &datasource)
		if err != nil {
			fmt.Fprintln(cmd.out, err)
		}

		ds, err := cmd.client.CreateDatasource(datasource)
		if err != nil {
			fmt.Fprintln(cmd.out, err)
		} else {
			fmt.Fprintln(cmd.out, "Datasource \""+ds.Name+"\" created")
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
