package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"

	"io/ioutil"

	"github.com/dimitrovvlado/grafctl/environment"
	"github.com/spf13/cobra"
)

type configCmd struct {
	out      io.Writer
	home     environment.Home
	hostname string
	username string
	password string
}

func newConfigCommand(out io.Writer) *cobra.Command {
	i := configCmd{out: out}

	configCmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"conf"},
		Short:   "Configure this command line tool",
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return i.run()
		},
	}
	f := configCmd.Flags()
	f.StringVarP(&i.hostname, "host", "", "", "Sets the grafana host")
	f.StringVarP(&i.username, "username", "u", "", "Sets the grafana username")
	f.StringVarP(&i.password, "password", "p", "", "Sets the grafana password")
	configCmd.MarkFlagRequired("host")
	configCmd.MarkFlagRequired("username")
	configCmd.MarkFlagRequired("password")
	return configCmd
}

func (i *configCmd) run() error {
	//TODO make home folder configurable via env. variables
	i.home = environment.Home(environment.DefaultHome)
	if err := ensureDirectories(i.home, i.out); err != nil {
		return err
	}

	log.Printf("grafctl has been configured at %s", i.home.String())

	var conf = map[string]string{"hostname": i.hostname, "username": i.username, "password": i.password}
	j, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(i.home.ConfigFile(), j, 0755); err != nil {
		return err
	}

	return nil
}

func ensureDirectories(home environment.Home, out io.Writer) error {
	configDirectories := []string{
		home.String(),
	}
	for _, p := range configDirectories {
		if fi, err := os.Stat(p); err != nil {
			log.Printf("Creating %s\n", p)
			if err := os.MkdirAll(p, 0755); err != nil {
				return fmt.Errorf("Could not create %s: %s", p, err)
			}
		} else if !fi.IsDir() {
			return fmt.Errorf("%s must be a directory", p)
		}
	}

	return nil
}

func ensureClient(client *grafana.Client) {
	if client == nil {
		log.Fatalln("Please configure grafctl with `grafctl config`")
	}
}
