package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dimitrovvlado/grafctl/grafana"
)

func TestListEmptyDatasourcePlain(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.DatasourcesEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write([]byte("[]"))
			},
		},
	})

	var buf bytes.Buffer
	cmd := newDatasourceListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "No datasources found", strings.TrimSpace(buf.String()))
}

func TestListEmptyDatasourceJson(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.DatasourcesEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write([]byte("[]"))
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"--output", "json"}
	cmd := newDatasourceListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "[]", strings.TrimSpace(buf.String()))
}

func TestListDatasourcesJson(t *testing.T) {
	dsBytes := helperLoadBytes(t, "datasources.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.DatasourcesEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(dsBytes)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"--output", "json"}
	cmd := newDatasourceListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})

	var localJSON interface{}
	json.Unmarshal(dsBytes, &localJSON)

	var remoteJSON interface{}
	json.Unmarshal(buf.Bytes(), &remoteJSON)

	require.True(t, reflect.DeepEqual(localJSON, remoteJSON))
}

func TestListDatasourcesPlain(t *testing.T) {
	dsBytes := helperLoadBytes(t, "datasources.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.DatasourcesEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(dsBytes)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newDatasourceListCommand(client, &buf)
	cmd.RunE(cmd, []string{})

	require.Contains(t, buf.String(), "1 	Prometheus	prometheus	proxy 	http://prometheus-server")
}

func TestGetDatasourceById(t *testing.T) {
	dsBytes := helperLoadBytes(t, "datasource.json")
	client := mockClient([]requestCase{
		{
			requestURI: fmt.Sprintf("%s/%s", grafana.DatasourcesEndpoint, "1"),
			handler: func(w http.ResponseWriter) {
				w.Write(dsBytes)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newDatasourceListCommand(client, &buf)
	cmd.RunE(cmd, []string{"1"})
	require.Contains(t, strings.TrimSpace(buf.String()), "1 	Prometheus	prometheus	proxy 	http://prometheus-server")
}
