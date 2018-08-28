package environment

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// EnvSettings describes all of the environment settings.
type EnvSettings struct {
	//Full URL of Grafana host
	GrafanaHost string
	//Personal username
	GrafanaUsername string
	//Grafana password
	GrafanaPassword string
}

// DefaultHome is the default folder of the configuration.
var DefaultHome = filepath.Join(homeDir(), ".config", "grafctl")

//Init environment
func (s *EnvSettings) Init() error {
	//TODO make home folder configurable via env. variables
	home := Home(DefaultHome)
	cfg := home.ConfigFile()

	b, err := ioutil.ReadFile(cfg)
	if err != nil {
		return err
	}

	jsonMap := make(map[string]string)
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}

	s.GrafanaHost = jsonMap["hostname"]
	s.GrafanaUsername = jsonMap["username"]
	s.GrafanaPassword = jsonMap["password"]

	return nil
}
