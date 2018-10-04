package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestDeleteFolder(t *testing.T) {
	b := helperLoadBytes(t, "folder.json")
	client := mockClient([]requestCase{
		{
			requestURI: fmt.Sprintf("%s/QOWKyoKmz", grafana.FoldersEndpoint),
			method:     http.MethodGet,
			handler: func(w http.ResponseWriter) {
				w.Write(b)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newFolderDeleteCmd(client, &buf)
	cmd.RunE(cmd, []string{"QOWKyoKmz"})

	require.Equal(t, "Folder \"Folder 1\" deleted\n", buf.String())
}
