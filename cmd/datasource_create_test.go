package cmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/stretchr/testify/require"
)

func TestCreateDatasource(t *testing.T) {
	dsBytes := helperLoadBytes(t, "datasource.json")
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
	flags := strings.Split("-f testdata/datasource.json", " ")
	cmd := newDatasourceCreateCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, flags)

	require.Equal(t, "Datasource \"Prometheus\" created\n", buf.String())
}
