package main

// EnvSettings describes all of the environment settings.
type EnvSettings struct {
	//Full URL of Grafana host
	GrafanaHost string
	//Personal username
	GrafanaUsername string
	//Grafana password
	GrafanaPassword string
}

//Init environment
func (s *EnvSettings) Init() {
	s.GrafanaHost = ""
	s.GrafanaUsername = ""
	s.GrafanaPassword = ""
}
