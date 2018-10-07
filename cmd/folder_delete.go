package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"
)

type deleteFolderCmd struct {
	client *grafana.Client
	out    io.Writer
	output string
	uid    string
}

func newFolderDeleteCmd(client *grafana.Client, out io.Writer) *cobra.Command {
	delete := &deleteFolderCmd{
		client: client,
		out:    out,
	}

	deleteFoldersCmd := &cobra.Command{
		Use:     "folder",
		Aliases: []string{"folders"},
		Short:   "Delete folder by UID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				log.Println("Command 'delete' requires an UID")
				cmd.Help()
				return nil
			}
			ensureClient(delete.client)
			delete.uid = args[0]
			return delete.run()
		},
	}

	return deleteFoldersCmd
}

func (i *deleteFolderCmd) run() error {
	folder, err := i.client.GetFolder(i.uid)
	if err != nil {
		switch err {
		case grafana.ErrNotFound:
			log.Printf("Folder with UID \"%s\" not found", i.uid)
		default:
			log.Println(err)
		}
	} else {
		err = i.client.DeleteFolder(folder)
		if err != nil {
			return err
		}
		fmt.Fprintln(i.out, fmt.Sprintf("Folder \"%s\" deleted", folder.Title))
	}
	return nil
}
