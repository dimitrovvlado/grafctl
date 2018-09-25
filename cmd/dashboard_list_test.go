package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestListDashboardsEmpty(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.SearchEndpoint,
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
			requestURI: grafana.SearchEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(dbBytes)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newDashboardListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Contains(t, buf.String(), "1 	cIBgcSjkk	Production Overview	db/production-overview	prod")
	require.NotContains(t, buf.String(), "163	Folder             	db/folder")
}

func TestListDashboardByUid(t *testing.T) {
	dbBytes := helperLoadBytes(t, "dashboardExport.json")
	client := mockClient([]requestCase{
		{
			requestURI: fmt.Sprintf("%s/QOWKyoKmz", grafana.DashboardsUIDEndpoint),
			handler: func(w http.ResponseWriter) {
				w.Write(dbBytes)
			},
		},
	})

	var buf bytes.Buffer
	cmd := newDashboardListCommand(client, &buf)
	cmd.RunE(cmd, []string{"QOWKyoKmz"})
	require.Contains(t, buf.String(), "database-backups	4      	General	/d/QOWKyoKmz/database-backups")
}
