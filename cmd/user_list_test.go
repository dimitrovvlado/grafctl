package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dimitrovvlado/grafctl/grafana"
)

func TestListEmptyUsersPlain(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.UsersEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write([]byte("[]"))
			},
		},
	})

	var buf bytes.Buffer
	cmd := newUsersListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "No users found", strings.TrimSpace(buf.String()))
}

func TestListEmptyUsersJson(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.UsersEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write([]byte("[]"))
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"--output", "json"}
	cmd := newUsersListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "[]", strings.TrimSpace(buf.String()))
}

func TestListUsersJson(t *testing.T) {
	userBytes := helperLoadBytes(t, "users.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.UsersEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(userBytes)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"--output", "json"}
	cmd := newUsersListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})

	var localJSON interface{}
	json.Unmarshal(userBytes, &localJSON)

	var remoteJSON interface{}
	json.Unmarshal(buf.Bytes(), &remoteJSON)

	require.True(t, reflect.DeepEqual(localJSON, remoteJSON))
}

func TestListUsersPlain(t *testing.T) {
	userBytes := helperLoadBytes(t, "users.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.UsersEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(userBytes)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newUsersListCommand(client, &buf)
	cmd.RunE(cmd, []string{})

	require.Contains(t, buf.String(), "1 	User Name	username	username@localhost")
}
