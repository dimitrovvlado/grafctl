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

type teamCreateCmd struct {
	client *grafana.Client
	out    io.Writer
	name   string
	email  string
	files  *[]string
}

func newTeamCreateCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	i := &teamCreateCmd{
		client: client,
		out:    out,
	}
	createTeamCmd := &cobra.Command{
		Use:     "team",
		Aliases: []string{},
		Short:   "Create team",
		Long:    "Create a team by either providing a file, or name and email",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(*i.files) == 0 && i.name == "" && i.email == "" {
				cmd.Help()
				os.Exit(0)
			}
			ensureClient(i.client)
			//TODO maybe change the command to create a team by a given set ot flags
			return i.run()
		},
	}

	i.files = createTeamCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename(s) or direcory to use to create the datasource")
	createTeamCmd.Flags().StringVarP(&i.name, "name", "n", "", "Set a team name")
	createTeamCmd.Flags().StringVarP(&i.email, "email", "e", "", "Set a team email")
	return createTeamCmd
}

func (i *teamCreateCmd) run() error {
	if len(*i.files) == 0 {
		if i.name == "" {
			return errors.New("Team email required")
		}
		if i.email == "" {
			return errors.New("Team name required")
		}
		_, err := i.client.CreateTeam(grafana.Team{Name: i.name, Email: i.email})
		if err != nil {
			switch err {
			case grafana.ErrConflict:
				log.Printf("Team \"%s\" already exists", i.name)
			default:
				log.Println(err)
			}
		} else {
			fmt.Fprintln(i.out, fmt.Sprintf("Team \"%s\" created", i.name))
		}
	} else {
		for _, file := range *i.files {
			importTeam(file, i)
		}
	}
	return nil
}

func importTeam(filename string, cmd *teamCreateCmd) {
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
				importTeam(fi, cmd)
			}
		}
	} else {
		//Ignoring the error because we check for file existance upfront
		data, _ := ioutil.ReadFile(filename)

		var team grafana.Team
		err := json.Unmarshal(data, &team)
		if err != nil {
			log.Printf("Error parsing %s", filename)
			return
		}

		_, err = cmd.client.CreateTeam(team)
		if err != nil {
			switch err {
			case grafana.ErrConflict:
				log.Printf("Team \"%s\" already exists", team.Name)
			default:
				log.Println(err)
			}
		} else {
			fmt.Fprintln(cmd.out, "Team \""+team.Name+"\" created")
		}
	}
}
