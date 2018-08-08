package cmd

import (
	"fmt"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

type usersCmd struct {
	client *grafana.Client
}

func newUsersCommand(client *grafana.Client) *cobra.Command {
	i := &usersCmd{
		client: client,
	}
	getUsersCmd := &cobra.Command{
		Use:     "users",
		Aliases: []string{"user"},
		Short:   "Display one or many users",
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return i.run()
		},
	}
	return getUsersCmd
}

// run creates a merge request
func (i *usersCmd) run() error {
	users, err := i.client.ListUsers()
	if err != nil {
		logrus.Fatal(err)
	}
	for _, u := range users {
		fmt.Printf("User: %s\n", u.Email)
	}

	return nil
}
