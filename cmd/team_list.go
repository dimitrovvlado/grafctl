package cmd

import (
	"fmt"
	"io"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type teamsCmd struct {
	client *grafana.Client
	out    io.Writer
	output string
	query  string
}

func newTeamsListCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	get := &teamsCmd{
		client: client,
		out:    out,
	}
	getTeamsCmd := &cobra.Command{
		Use:     "teams",
		Aliases: []string{"team"},
		Short:   "Display one or many teams",
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(get.client)
			return get.run()
		},
	}
	f := getTeamsCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
	f.StringVarP(&get.query, "query", "q", "", "Query string")
	return getTeamsCmd
}

func (i *teamsCmd) run() error {
	teamsPage, err := i.client.SearchTeams(&grafana.SearchTeamsOptions{Query: i.query})
	if err != nil {
		logrus.Fatal(err)
	}
	teams := teamsPage.Teams
	//TODO extract as flag
	var colWidth uint = 60
	formatter := func() string {
		if teams == nil || len(teams) == 0 {
			return fmt.Sprintf("No teams found.")
		}
		table := uitable.New()
		table.MaxColWidth = colWidth
		table.AddRow("ID", "NAME", "EMAIL", "MEMBERS")
		for _, lr := range teams {
			table.AddRow(lr.ID, lr.Name, lr.Email, lr.MemberCount)
		}
		return fmt.Sprintf("%s%s", table.String(), "\n")
	}

	result, err := formatResult(i.output, teams, formatter)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Fprintln(i.out, result)
	return nil
}
