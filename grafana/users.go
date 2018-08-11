package grafana

import (
	"fmt"
	"net/http"
)

// ListUsers returns a matched courier based on tracking number.
func (c *Client) ListUsers() ([]User, error) {
	resp, err := c.doRequest(
		http.MethodGet,
		UsersEndpoint,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Decode the response into a AfterShip response object.
	var users []User
	if err := decodeResponse(resp, &users); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return nil, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	// Check if we didn't get a result and return an error if true.
	if users == nil || len(users) <= 0 {
		return make([]User, 0), nil
	}

	return users, nil
}

// ListOrgs returns a matched courier based on tracking number.
func (c *Client) ListOrgs() ([]Org, error) {
	resp, err := c.doRequest(
		http.MethodGet,
		UsersEndpoint,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Decode the response into a AfterShip response object.
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
