package cmd

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dimitrovvlado/grafctl/grafana"
)

func TestListEmptyOrgsPlain(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.OrgsEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write([]byte("[]"))
			},
		},
	})

	var buf bytes.Buffer
	cmd := newOrgListCommand(client, &buf)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "No organizations found.", strings.TrimSpace(buf.String()))
}

func TestListEmptyOrgsJson(t *testing.T) {
	client := mockClient([]requestCase{
		{
			requestURI: grafana.OrgsEndpoint,
			handler: func(w http.ResponseWriter) {
				w.Write([]byte("[]"))
			},
		},
	})

	var buf bytes.Buffer
	flags := []string{"--output", "json"}
	cmd := newOrgListCommand(client, &buf)
	cmd.ParseFlags(flags)
	cmd.RunE(cmd, []string{})
	require.Equal(t, "[]", strings.TrimSpace(buf.String()))
}
