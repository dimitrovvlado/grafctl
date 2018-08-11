package cmd

import (
	"fmt"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

type orgCmd struct {
	client *grafana.Client
	output string
}

func newOrgCommand(client *grafana.Client) *cobra.Command {
	get := &orgCmd{
		client: client,
	}
	getOrgsCmd := &cobra.Command{
		Use:     "organinizations",
		Aliases: []string{"organisations", "orgs", "org"},
		Short:   "Display one or many organizations",
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			get.output = cmd.Flag("output").Value.String()
			return get.run()
		},
	}
	return getOrgsCmd
}

// run creates a merge request
func (i *orgCmd) run() error {
	orgs, err := i.client.ListOrgs()
	if err != nil {
		logrus.Fatal(err)
	}

	//TODO extract as flag
	var colWidth uint = 60
	formatter := func() string {
		table := uitable.New()
		table.MaxColWidth = colWidth
		table.AddRow("ID", "NAME", "CITY", "STATE", "COUNTRY")
		for _, lr := range orgs {
			table.AddRow(lr.ID, lr.Name, lr.Address.City, lr.Address.State, lr.Address.Country)
		}
		return fmt.Sprintf("%s%s", table.String(), "\n")
	}

	result, err := formatResult(i.output, orgs, formatter)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf(result)

	return nil
}
