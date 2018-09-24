package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"
)

type dashboardCreateCmd struct {
	client    *grafana.Client
	out       io.Writer
	files     *[]string
	overwrite bool
	message   string
}

func newDashboardCreateCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	i := &dashboardCreateCmd{
		client: client,
		out:    out,
	}
	createDashboardCmd := &cobra.Command{
		Use:     "dashboard",
		Aliases: []string{"db"},
		Short:   "Create dashboard",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(i.client)
			return i.run()
		},
	}

	i.files = createDashboardCmd.PersistentFlags().StringSliceP("filename", "f", []string{}, "Filename(s) or direcory to use to create the dashboards")
	createDashboardCmd.MarkPersistentFlagRequired("filename")
	createDashboardCmd.Flags().StringVarP(&i.message, "message", "m", "", "Set a commit message for the version history")
	createDashboardCmd.Flags().BoolVarP(&i.overwrite, "overwrite", "w", false, "Overwrite existing dashboard with newer version")
	return createDashboardCmd
}

func (i *dashboardCreateCmd) run() error {
	// dbs, err := i.client.ListDashboards()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// slugs := make(map[string]int)
	// for _, db := range dbs {
	// 	slugs[db.URI[3:]] = db.ID
	// }

	for _, file := range *i.files {
		importDаshboard(file, i)
	}
	return nil
}

func importDаshboard(filename string, cmd *dashboardCreateCmd) {
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
				importDаshboard(fi, cmd)
			}
		}
	} else {
		//Ignoring the error because we check for file existance upfront
		data, _ := ioutil.ReadFile(filename)

		var dashboard map[string]interface{}
		err := json.Unmarshal(data, &dashboard)
		if err != nil {
			log.Printf("Error parsing %s", filename)
			return
		}
		dashboard["id"] = nil

		dbr := grafana.DashboardRequest{
			Message:   cmd.message,
			Overwrite: &cmd.overwrite, //false values included
		}
		dbr.Dashboard = dashboard
		dbr.Message = cmd.message

		dr, err := cmd.client.CreateDashboard(dbr)
		if err != nil {
			switch err.(type) {
			case *grafana.Error:
				//Grafana returns error code 500 if dashboard cannot be imported
				fmt.Printf("%s: %s\n", err.Error(), dashboard["title"])
			default:
				//For all other
				log.Fatalln(err)
			}
		} else {
			fmt.Fprintln(cmd.out, "Dashboard \""+dr.Title+"\" imported")
		}
	}
}
