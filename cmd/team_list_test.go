package cmd

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dimitrovvlado/grafctl/grafana"
)

func TestListEmptyTeamPlain(t *testing.T) {
	teamBytes := helperLoadBytes(t, "emptyTeams.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.TeamsEndpoint + "?query=",
			handler: func(w http.ResponseWriter) {
				w.Write(teamBytes)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newTeamsListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "No teams found.", strings.TrimSpace(buf.String()))
}

func TestListEmptyTeamJson(t *testing.T) {
	teamBytes := helperLoadBytes(t, "emptyTeams.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.TeamsEndpoint + "?query=",
			handler: func(w http.ResponseWriter) {
				w.Write(teamBytes)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"--output", "json"}
	cmd := newTeamsListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "[]", strings.TrimSpace(buf.String()))
}
func TestListTeamsPlain(t *testing.T) {
	teamBytes := helperLoadBytes(t, "teams.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.TeamsEndpoint + "?query=",
			handler: func(w http.ResponseWriter) {
				w.Write(teamBytes)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newTeamsListCommand(client, &buf)
	cmd.RunE(cmd, []string{})

	require.Contains(t, buf.String(), "1 	TestTeam	team@localhost	1")
}
