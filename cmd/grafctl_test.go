package cmd

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/dimitrovvlado/grafctl/grafana"
)

type requestCase struct {
	method     string
	requestURI string
	handler    func(w http.ResponseWriter)
}

func mockClient(cases []requestCase) *grafana.Client {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, req := range cases {
			if req.requestURI == r.RequestURI {
				if req.method != "" && req.method == r.Method {
					req.handler(w)
				} else {
					req.handler(w)
				}
			}
		}
	}))
	return grafana.New(apiStub.URL, "username", "password")
}

func helperLoadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}
