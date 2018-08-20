package cmd

import (
	"fmt"
	"io"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

type orgCmd struct {
	client *grafana.Client
	out    io.Writer
	output string
}

func newOrgListCommand(client *grafana.Client, out io.Writer) *cobra.Command {
	get := &orgCmd{
		client: client,
		out:    out,
	}
	getOrgsCmd := &cobra.Command{
		Use:     "organinizations",
		Aliases: []string{"organisations", "orgs", "org"},
		Short:   "Display the current organization.",
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return get.run()
		},
	}
	f := getOrgsCmd.Flags()
	f.StringVarP(&get.output, "output", "o", "", "Output the specified format (|json)")
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
	fmt.Fprintln(i.out, result)

	return nil
}
