package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestListUsersJson(t *testing.T) {
	userBytes := helperLoadBytes(t, "users.json")
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case grafana.UsersEndpoint:
			w.Write(userBytes)
		default:
			return
		}
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	flags := strings.Split("--output json", " ")
	cmd := newUsersListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, flags)

	var localJSON interface{}
	json.Unmarshal(userBytes, &localJSON)

	var remoteJSON interface{}
	json.Unmarshal(buf.Bytes(), &remoteJSON)

	require.True(t, reflect.DeepEqual(localJSON, remoteJSON))
}

func TestListUsersPlain(t *testing.T) {
	userBytes := helperLoadBytes(t, "users.json")
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case grafana.UsersEndpoint:
			w.Write(userBytes)
		default:
			return
		}
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	cmd := newUsersListCommand(client, &buf)
	cmd.RunE(cmd, []string{})

	require.Contains(t, buf.String(), "1 	User Name	username	username@localhost")
}
