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

func TestListEmptyUsersPlain(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp string
		switch r.RequestURI {
		case grafana.UsersEndpoint:
			resp = "[]"
		default:
			return
		}
		w.Write([]byte(resp))
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	cmd := newUsersListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "ID	NAME	LOGIN	EMAIL", strings.TrimSpace(buf.String()))
}

func TestListEmptyUsersJson(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp string
		switch r.RequestURI {
		case grafana.UsersEndpoint:
			resp = "[]"
		default:
			return
		}
		w.Write([]byte(resp))
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	flags := strings.Split("--output json", " ")
	cmd := newUsersListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, flags)
	require.Equal(t, "[]", strings.TrimSpace(buf.String()))
}
