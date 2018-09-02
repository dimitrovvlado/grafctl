package cmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dimitrovvlado/grafctl/grafana"
)

func TestListEmptyTeamPlain(t *testing.T) {
	teamBytes := helperLoadBytes(t, "emptyTeams.json")
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case grafana.TeamsEndpoint + "?query=":
			w.Write(teamBytes)
		default:
			return
		}
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	cmd := newTeamsListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "No teams found.", strings.TrimSpace(buf.String()))
}

func TestListEmptyTeamJson(t *testing.T) {
	teamBytes := helperLoadBytes(t, "emptyTeams.json")
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case grafana.TeamsEndpoint + "?query=":
			w.Write(teamBytes)
		default:
			return
		}
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	flags := strings.Split("--output json", " ")
	cmd := newTeamsListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, flags)
	require.Equal(t, "[]", strings.TrimSpace(buf.String()))
}
func TestListTeamsPlain(t *testing.T) {
	teamBytes := helperLoadBytes(t, "teams.json")
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case grafana.TeamsEndpoint + "?query=":
			w.Write(teamBytes)
		default:
			return
		}
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	cmd := newTeamsListCommand(client, &buf)
	cmd.RunE(cmd, []string{})

	require.Contains(t, buf.String(), "1 	TestTeam	team@localhost	1")
}
