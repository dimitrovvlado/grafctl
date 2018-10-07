package cmd

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestCreateFolder(t *testing.T) {
	dsBytes := helperLoadBytes(t, "folder.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.FoldersEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(dsBytes)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"-f", "testdata/folder.json"}
	cmd := newFolderCreateCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})

	require.Equal(t, "Folder \"Folder 1\" created\n", buf.String())
}
