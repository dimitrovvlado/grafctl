package cmd

import (
	"fmt"
	"io"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

type usersCmd struct {
	client     *grafana.Client
	out        io.Writer
	output     string
	currentOrg bool
}

func newUsersListCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	get := &usersCmd{
		client: client,
		out:    out,
	}
	getUsersCmd := &cobra.Command{
		Use:     "users",
		Aliases: []string{"user"},
		Short:   "Display one or many users",
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return get.run()
		},
	}
	f := getUsersCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
	getUsersCmd.PersistentFlags().BoolVarP(&get.currentOrg, "current-org", "c", false, "Display users in current organization only")
	return getUsersCmd
}

// run creates a merge request
func (i *usersCmd) run() error {
	users, err := i.client.ListUsers(&grafana.ListUserOptions{CurrentOrg: i.currentOrg})
	if err != nil {
		logrus.Fatal(err)
	}

	//TODO extract as flag
	var colWidth uint = 60
	formatter := func() string {
		table := uitable.New()
		table.MaxColWidth = colWidth
		table.AddRow("ID", "NAME", "LOGIN", "EMAIL")
		for _, lr := range users {
			userID := lr.ID
			if userID == 0 {
				userID = lr.UserID
			}
			table.AddRow(userID, lr.Name, lr.Login, lr.Email)
		}
		return fmt.Sprintf("%s%s", table.String(), "\n")
	}

	result, err := formatResult(i.output, users, formatter)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Fprintln(i.out, result)
	return nil
}
