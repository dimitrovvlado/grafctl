package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"
)

type teamDeleteCmd struct {
	client *grafana.Client
	out    io.Writer
	teamID string
}

func newTeamDeleteCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	i := &teamDeleteCmd{
		client: client,
		out:    out,
	}
	deleteTeamCmd := &cobra.Command{
		Use:     "team",
		Aliases: []string{},
		Short:   "Delete team by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				log.Println("Command 'delete' requires an ID")
				cmd.Help()
				return nil
			}
			ensureClient(i.client)
			i.teamID = args[0]
			return i.run()
		},
	}

	return deleteTeamCmd
}

// run deletes a ds
func (i *teamDeleteCmd) run() error {
	team, err := i.client.GetTeam(i.teamID)
	if err != nil {
		switch err {
		case grafana.ErrNotFound:
			log.Printf("Team with ID \"%s\" not found", i.teamID)
		default:
			log.Println(err)
		}
	} else {
		err = i.client.DeleteTeam(team)
		if err != nil {
			return err
		}
		fmt.Fprintln(i.out, fmt.Sprintf("Team \"%s\" deleted", team.Name))
	}
	return nil
}
