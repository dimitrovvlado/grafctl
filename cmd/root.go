package cmd

import (
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"

	"github.com/spf13/cobra"
)

type rootCmd struct {
	Verbous bool
}

//NewRootCmd creates the root command
func NewRootCmd(client *grafana.Client) *cobra.Command {
	i := &rootCmd{}

	rootCmd := &cobra.Command{
		Use:   "grafctl",
		Short: "Grafctl is command line tool for managing Grafana",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&i.Verbous, "verbose", "v", false, "Verbose output")

	rootCmd.AddCommand(
		newVersionCmd(),
		newGetCmd(client))

	return rootCmd
}

// func newRootCmd(args []string) *cobra.Command {

// 	cmd := &cobra.Command{
// 		Use:          "grafctl",
// 		Short:        "The Grafana management tool.",
// 		Long:         globalUsage,
// 		SilenceUsage: true,
// 	}
// 	flags := cmd.PersistentFlags()
// 	flags.Parse(args)

// 	out := cmd.OutOrStdout()

// 	cmd.AddCommand(
// 		newGetCommand(out),
// 	)

// 	settings.Init()

// 	return cmd
// }
