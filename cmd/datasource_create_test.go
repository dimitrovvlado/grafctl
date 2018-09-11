package cmd

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestCreateDatasource(t *testing.T) {
	dsBytes := helperLoadBytes(t, "datasource.json")
	client := mockClient([]requestCase{
		{
			requestURI: grafana.DatasourcesEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write(dsBytes)
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"-f", "testdata/datasource.json"}
	cmd := newDatasourceCreateCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, flags)

	require.Equal(t, "Datasource \"Prometheus\" created\n", buf.String())
}
