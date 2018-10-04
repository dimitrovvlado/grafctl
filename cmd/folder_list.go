package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

type getFolderCmd struct {
	client *grafana.Client
	out    io.Writer
	output string
	uid    string
}

func newFoldersListCmd(client *grafana.Client, out io.Writer) *cobra.Command {
	get := &getFolderCmd{
		client: client,
		out:    out,
	}

	getFoldersCmd := &cobra.Command{
		Use:     "folders",
		Aliases: []string{"folder"},
		Short:   "Display one or many folders",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(get.client)
			if len(args) > 0 {
				get.uid = args[0]
			}
			return get.run()
		},
	}

	f := getFoldersCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
	return getFoldersCmd
}

func (i *getFolderCmd) run() error {
	//TODO extract as flag
	var colWidth uint = 60
	var formatter func() string
	var obj interface{}
	if i.uid != "" {
		folder, err := i.client.GetFolder(i.uid)
		if err != nil {
			log.Fatalln(err)
		}
		formatter = func() string {
			if (grafana.Folder{}) == folder {
				return fmt.Sprintf("Folder not found")
			}
			table := uitable.New()
			table.MaxColWidth = colWidth
			table.AddRow("ID", "UID", "TITLE", "URL")
			table.AddRow(folder.ID, folder.UID, folder.Title, folder.URL)
			return fmt.Sprintf("%s%s", table.String(), "\n")
		}
		obj = folder
	} else {
		folders, err := i.client.ListFolders()
		if err != nil {
			log.Fatalln(err)
		}
		formatter = func() string {
			if folders == nil || len(folders) == 0 {
				return fmt.Sprintf("No folders found")
			}
			table := uitable.New()
			table.MaxColWidth = colWidth
			table.AddRow("ID", "UID", "TITILE")
			for _, lf := range folders {
				table.AddRow(lf.ID, lf.UID, lf.Title)
			}
			return fmt.Sprintf("%s%s", table.String(), "\n")
		}
		obj = folders
	}

	result, err := formatResult(i.output, obj, formatter)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(i.out, result)
	return nil
}
