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
	s.GrafanaHost = "https://vcpp1-dev.us-west-2.csp.vmware.com"
	s.GrafanaUsername = "usagemeter"
	s.GrafanaPassword = "axDaVUat0oyeeJx5AIDVfFZTxIwR"
}
