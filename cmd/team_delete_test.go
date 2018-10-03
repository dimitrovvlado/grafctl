package cmd

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestDeleteTeam(t *testing.T) {
	b := helperLoadBytes(t, "team.json")
	client := mockClient([]requestCase{
		{
			requestURI: fmt.Sprintf("%s/3", grafana.TeamsEndpoint),
			method:     http.MethodGet,
			handler: func(w http.ResponseWriter) {
				w.Write(b)
			},
		},
		{
			requestURI: fmt.Sprintf("%s/3", grafana.DashboardsUIDEndpoint),
			method:     http.MethodDelete,
			handler: func(w http.ResponseWriter) {
				w.Write(b)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newTeamDeleteCommand(client, &buf)
	cmd.RunE(cmd, []string{"3"})

	require.Equal(t, "Team \"MyTestTeam\" deleted\n", buf.String())
}

func TestDeleteNotExistingTeam(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: fmt.Sprintf("%s/3", grafana.TeamsEndpoint),
			method:     http.MethodGet,
			handler: func(w http.ResponseWriter) {
				w.WriteHeader(404)
			},
		},
	})

	var buf bytes.Buffer
	log.SetOutput(&buf)
	cmd := newTeamDeleteCommand(client, &buf)
	cmd.RunE(cmd, []string{"3"})
	require.Equal(t, "Team with ID \"3\" not found\n", buf.String())
}

func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}
