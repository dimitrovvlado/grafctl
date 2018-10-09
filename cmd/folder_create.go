package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"
)

type folderCreateCmd struct {
	client *grafana.Client
	out    io.Writer
	title  string
	files  *[]string
}

func newFolderCreateCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	i := &folderCreateCmd{
		client: client,
		out:    out,
	}

	createFolderCmd := &cobra.Command{
		Use:     "folder",
		Aliases: []string{"folders"},
		Short:   "Create folder",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(i.client)
			return i.run()
		},
	}

	createFolderCmd.Flags().StringVarP(&i.title, "title", "t", "", "Set a title of the folder")
	i.files = createFolderCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename(s) or direcory to use to create the remote folders")
	return createFolderCmd
}

func (i *folderCreateCmd) run() error {
	if len(*i.files) == 0 {
		if i.title == "" {
			return errors.New("Either folder name or file location is required")
		}
		f, err := i.client.CreateFolder(grafana.Folder{Title: i.title})
		if err != nil {
			switch err.(type) {
			case *grafana.Error:
				fmt.Printf("%s: %s\n", err.Error(), i.title)
			default:
				//For all other
				log.Fatalln(err)
			}
		} else {
			fmt.Fprintln(i.out, "Folder \""+f.Title+"\" created")
		}
	} else {
		for _, file := range *i.files {
			importFolder(file, i)
		}
	}
	return nil
}

func importFolder(filename string, cmd *folderCreateCmd) {
	info, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
		return
	}
	if info.IsDir() {
		files, err := filePathWalkDir(filename)
		if err != nil {
			log.Println(err)
		} else {
			for _, fi := range files {
				importFolder(fi, cmd)
			}
		}
	} else {
		//Ignoring the error because we check for file existence upfront
		data, _ := ioutil.ReadFile(filename)

		var folder grafana.Folder
		err := json.Unmarshal(data, &folder)
		if err != nil {
			log.Printf("Error parsing %s", filename)
			return
		}

		f, err := cmd.client.CreateFolder(folder)
		if err != nil {
			switch err.(type) {
			case *grafana.Error:
				fmt.Println(err.Error())
			default:
				//For all other
				log.Fatalln(err)
			}
		} else {
			fmt.Fprintln(cmd.out, "Folder \""+f.Title+"\" created")
		}
	}
}
