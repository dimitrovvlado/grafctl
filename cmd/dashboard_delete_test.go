package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestDeleteDashboard(t *testing.T) {
	b := []byte("{\"title\": \"UI\"}")
	client := mockClient([]requestCase{
		{
			requestURI: fmt.Sprintf("%s/QOWKyoKmz", grafana.DashboardsUIDEndpoint),
			handler: func(w http.ResponseWriter) {
				w.Write(b)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"QOWKyoKmz"}
	cmd := newDashboardDeleteCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, flags)

	require.Equal(t, "Dashboard \"UI\" deleted\n", buf.String())
}
