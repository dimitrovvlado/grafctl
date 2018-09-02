package grafana

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)

//Client for http requests
type Client struct {
	apiURI   string
	username string
	password string
	verbose  bool
}

// New creates a new Grafana API client.
func New(apiURI, username, password string) *Client {
	return &Client{
		apiURI:   apiURI,
		username: username,
		password: password,
	}
}

//SetVerbose enables verbous logging of requests and responses
func (c *Client) SetVerbose(verbose bool) {
	c.verbose = verbose
}

type request struct {
	method   string
	endpoint string
	data     interface{}
	query    map[string]string
}

//func (c *Client) doRequest(method, endpoint string, data interface{}, query map[string]string) (*http.Response, error) {
func (c *Client) doRequest(req *request) (*http.Response, error) {
	client := http.DefaultClient

	// Encode data if we are passed an object.
	b := bytes.NewBuffer(nil)
	if req.data != nil {
		// Create the encoder.
		enc := json.NewEncoder(b)
		if err := enc.Encode(req.data); err != nil {
			return nil, fmt.Errorf("json encoding data for doRequest failed: %v", err)
		}
	}

	// Create the request.
	uri := fmt.Sprintf("%s/%s", c.apiURI, strings.Trim(req.endpoint, "/"))

	httpReq, err := http.NewRequest(req.method, uri, b)
	if err != nil {
		return nil, fmt.Errorf("creating %s request to %s failed: %v", req.method, uri, err)
	}

	// Set the correct headers.
	httpReq.Header.Set("Authorization", AuthorizationTypeBasic+basicAuth(c.username, c.password))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	// Set query parameters if any
	if req.query != nil {
		q := httpReq.URL.Query()
		for k, v := range req.query {
			q.Set(k, v)
		}
		httpReq.URL.RawQuery = q.Encode()
	}

	if c.verbose {
		debug, err := httputil.DumpRequestOut(httpReq, true)
		if err == nil {
			fmt.Println(string(debug))
		}
	}

	// Do the request.
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("performing %s request to %s failed: %v", req.method, uri, err)
	}

	if c.verbose {
		debug, err := httputil.DumpResponse(resp, true)
		if err == nil {
			fmt.Println(string(debug))
		}
	}

	// Check that the response status code was OK.
	if resp.StatusCode > 400 {
		// Read the body of the request, ignore the error since we are already in the error state.
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		// Create a friendly error message based off the status code returned.
		var message string
		switch resp.StatusCode {
		case http.StatusUnauthorized: // 401
			message = "Unauthorized."
		case http.StatusForbidden: // 403
			message = "The request is understood, but it has been refused or access is not allowed."
		case http.StatusNotFound: // 404
			message = "The URI requested is invalid or the resource requested does not exist."
		case http.StatusTooManyRequests: // 429
			message = "You have exceeded the API call rate limit. Default limit is 10 requests per second."
		case http.StatusInternalServerError: // 500
			message = "Something went wrong on Grafana's end."
		case http.StatusNotImplemented: // 501
			message = "Something went wrong on Grafana's end."
		case http.StatusBadGateway: // 502
			message = "Something went wrong on Grafana's end."
		case http.StatusServiceUnavailable: // 503
			message = "Something went wrong on Grafana's end."
		}

		return nil, fmt.Errorf("%s request to %s returned status code %d: message -> %s\nbody -> %s", req.method, uri, resp.StatusCode, message, string(body))
	}

	// Return errors on the API errors.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return resp, fmt.Errorf("API error")
	}

	return resp, nil
}

func decodeResponse(resp *http.Response, v interface{}) error {
	// Copy buffer so we have a backup.
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.Unmarshal(buf.Bytes(), v)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
