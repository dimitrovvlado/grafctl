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

func TestListEmptyDatasourcePlain(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case grafana.DatasourcesEndpoint:
			w.Write([]byte("[]"))
		default:
			return
		}
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	cmd := newDatasourceListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "ID	NAME	TYPE	ACCESS	URL", strings.TrimSpace(buf.String()))
}

func TestListEmptyDatasourceJson(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case grafana.DatasourcesEndpoint:
			w.Write([]byte("[]"))
		default:
			return
		}
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	flags := strings.Split("--output json", " ")
	cmd := newDatasourceListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, flags)
	require.Equal(t, "[]", strings.TrimSpace(buf.String()))
}

func TestListDatasourcesJson(t *testing.T) {
	dsBytes := helperLoadBytes(t, "datasources.json")
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case grafana.DatasourcesEndpoint:
			w.Write(dsBytes)
		default:
			return
		}
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	flags := strings.Split("--output json", " ")
	cmd := newDatasourceListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, flags)

	var localJSON interface{}
	json.Unmarshal(dsBytes, &localJSON)

	var remoteJSON interface{}
	json.Unmarshal(buf.Bytes(), &remoteJSON)

	require.True(t, reflect.DeepEqual(localJSON, remoteJSON))
}

func TestListDatasourcesPlain(t *testing.T) {
	dsBytes := helperLoadBytes(t, "datasources.json")
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case grafana.DatasourcesEndpoint:
			w.Write(dsBytes)
		default:
			return
		}
	}))

	client := grafana.New(apiStub.URL, "username", "password")

	var buf bytes.Buffer
	cmd := newDatasourceListCommand(client, &buf)
	cmd.RunE(cmd, []string{})

	require.Contains(t, buf.String(), "1 	Prometheus	prometheus	proxy 	http://prometheus-server")
}
