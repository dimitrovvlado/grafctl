package cmd

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestCreateTeamFromFile(t *testing.T) {
	tb := []byte("{\"message\":\"Team created\",\"teamId\":2}")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.TeamsEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(tb)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"-f", "testdata/team.json"}
	cmd := newTeamCreateCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})

	require.Equal(t, "Team \"MyTestTeam\" created\n", buf.String())
}
func TestCreateTeamFromFlags(t *testing.T) {
	tb := []byte("{\"message\":\"Team created\",\"teamId\":2}")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.TeamsEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(tb)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"-n", "TestTeam", "-e", "test@test.com"}
	cmd := newTeamCreateCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})

	require.Equal(t, "Team \"TestTeam\" created\n", buf.String())
}
