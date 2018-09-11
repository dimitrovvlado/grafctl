package grafana

import (
	"fmt"
	"net/http"
)

// ListOrgs returns a list of organizations.
func (c *Client) ListOrgs() ([]Org, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: OrgsEndpoint,
	})
	if err != nil {
		return nil, err
	}

	// Decode the response into a []Org response object.
	var orgs []Org
	if err := decodeResponse(resp, &orgs); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return nil, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	// Check if we didn't get a result and return an error if true.
	if orgs == nil || len(orgs) <= 0 {
		return make([]Org, 0), nil
	}

	return orgs, nil
}
