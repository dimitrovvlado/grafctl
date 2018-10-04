package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestListEmptyFoldersPlain(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.FoldersEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write([]byte("[]"))
			},
		},
	})

	var buf bytes.Buffer
	cmd := newFoldersListCmd(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "No folders found", strings.TrimSpace(buf.String()))
}

func TestListEmptyFoldersJson(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.FoldersEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write([]byte("[]"))
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"--output", "json"}
	cmd := newFoldersListCmd(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "[]", strings.TrimSpace(buf.String()))
}

func TestListFoldersJson(t *testing.T) {
	folderBytes := helperLoadBytes(t, "folders.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.FoldersEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(folderBytes)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"--output", "json"}
	cmd := newFoldersListCmd(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})

	var localJSON interface{}
	json.Unmarshal(folderBytes, &localJSON)

	var remoteJSON interface{}
	json.Unmarshal(buf.Bytes(), &remoteJSON)

	require.True(t, reflect.DeepEqual(localJSON, remoteJSON))
}

func TestGetFolderByIdPlain(t *testing.T) {
	folderBytes := helperLoadBytes(t, "folder.json")
	client := mockClient([]requestCase{
		{
			requestURI: fmt.Sprintf("%s/%s", grafana.FoldersEndpoint, "1"),
			handler: func(w http.ResponseWriter) {
				w.Write(folderBytes)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newFoldersListCmd(client, &buf)
	cmd.RunE(cmd, []string{"1"})
	require.Contains(t, strings.TrimSpace(buf.String()), "35	CLj2e60iz	Folder 1	/dashboards/f/CLj2e60iz/folder-1")
}

func TestGetFolderByIdJson(t *testing.T) {
	folderBytes := helperLoadBytes(t, "folder.json")
	client := mockClient([]requestCase{
		{
			requestURI: fmt.Sprintf("%s/%s", grafana.FoldersEndpoint, "1"),
			handler: func(w http.ResponseWriter) {
				w.Write(folderBytes)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"--output", "json"}
	cmd := newFoldersListCmd(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{"1"})

	var localJSON interface{}
	json.Unmarshal(folderBytes, &localJSON)

	var remoteJSON interface{}
	json.Unmarshal(buf.Bytes(), &remoteJSON)

	require.True(t, reflect.DeepEqual(localJSON, remoteJSON))
}
