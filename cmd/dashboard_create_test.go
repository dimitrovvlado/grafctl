package cmd

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestCreateDashboard(t *testing.T) {
	dsBytes := helperLoadBytes(t, "dashboardResponse.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.DashboardsImportEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(dsBytes)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"-f", "testdata/dashboard.json"}
	cmd := newDashboardCreateCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})

	require.Equal(t, "Dashboard \"UI\" imported\n", buf.String())
}
