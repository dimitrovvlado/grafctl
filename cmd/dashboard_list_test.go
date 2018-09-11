package cmd

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestListDashboardsEmpty(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.DashboardSearchEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write([]byte("[]"))
			},
		},
	})

	var buf bytes.Buffer
	cmd := newDashboardListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "No dashboards found.", strings.TrimSpace(buf.String()))
}

func TestListDashboardsFilter(t *testing.T) {
	dbBytes := helperLoadBytes(t, "dashboardsAndFolders.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.DashboardSearchEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(dbBytes)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newDashboardListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Contains(t, buf.String(), "1 	Production Overview	db/production-overview	prod")
	require.NotContains(t, buf.String(), "163	Folder             	db/folder")
}
