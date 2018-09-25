package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"
)

type usersCmd struct {
	client     *grafana.Client
	out        io.Writer
	output     string
	currentOrg bool
	ID         string
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
		RunE: func(cmd *cobra.Command, args []string) error {
			ensureClient(get.client)
			if len(args) > 0 {
				get.ID = args[0]
			}
			return get.run()
		},
	}
	f := getUsersCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
	f.BoolVarP(&get.currentOrg, "current-org", "c", false, "Display users in current organization only")
	return getUsersCmd
}

func (i *usersCmd) run() error {
	//TODO extract as flag
	var colWidth uint = 60
	var formatter func() string
	var obj interface{}
	if i.ID != "" {
		user, err := i.client.GetUser(i.ID)
		if err != nil {
			log.Fatalln(err)
		}
		formatter = func() string {
			if (grafana.User{}) == user {
				return fmt.Sprintf("User not found")
			}
			table := uitable.New()
			table.MaxColWidth = colWidth
			table.AddRow("ID", "NAME", "LOGIN", "EMAIL")
			table.AddRow(i.ID, user.Name, user.Login, user.Email)
			return fmt.Sprintf("%s%s", table.String(), "\n")
		}
		obj = user
	} else {
		users, err := i.client.ListUsers(&grafana.ListUserOptions{CurrentOrg: i.currentOrg})
		if err != nil {
			log.Fatalln(err)
		}
		formatter = func() string {
			if users == nil || len(users) == 0 {
				return fmt.Sprintf("No users found")
			}
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
		obj = users
	}

	result, err := formatResult(i.output, obj, formatter)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(i.out, result)
	return nil
}
