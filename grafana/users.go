package grafana

import (
	"fmt"
	"net/http"
)

//TODO rename file

// ListUsers returns a list of users.
func (c *Client) ListUsers(opt *ListUserOptions) ([]User, error) {
	var endpoint string
	if opt.CurrentOrg {
		endpoint = OrgsUsersEndpoint
	} else {
		endpoint = UsersEndpoint
	}
	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: endpoint,
	})
	if err != nil {
		return nil, err
	}

	// Decode the response into a []User response object.
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

// GetDatasource returns a datasource by ID
func (c *Client) GetDatasource(id string) (Datasource, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s", DatasourcesEndpoint, id),
	})
	if err != nil {
		return Datasource{}, err
	}

	// Decode the response into a Datasource response object.
	var ds Datasource
	if err := decodeResponse(resp, &ds); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return Datasource{}, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	return ds, nil
}

// ListDatasources returns a list of datasources
func (c *Client) ListDatasources() ([]Datasource, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: DatasourcesEndpoint,
	})
	if err != nil {
		return nil, err
	}

	// Decode the response into a []Datasource response object.
	var ds []Datasource
	if err := decodeResponse(resp, &ds); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return nil, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	// Check if we didn't get a result and return an error if true.
	if ds == nil || len(ds) <= 0 {
		return make([]Datasource, 0), nil
	}

	return ds, nil
}

// CreateDatasource creates a datasource
func (c *Client) CreateDatasource(ds Datasource) (Datasource, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodPost,
		endpoint: DatasourcesEndpoint,
		data:     ds,
	})
	// Decode the response into a Datasource response object.
	var res Datasource
	if err != nil {
		return res, err
	}

	if err := decodeResponse(resp, &res); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return res, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	return res, nil
}

// DeleteDatasource deletes a datasource
func (c *Client) DeleteDatasource(ds Datasource) error {
	resp, err := c.doRequest(&request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("%s/%d", DatasourcesEndpoint, ds.ID),
	})

	defer resp.Body.Close()
	return err
}

//SearchTeams returns a list of teams
func (c *Client) SearchTeams(opt *SearchTeamsOptions) (TeamPage, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: TeamsEndpoint,
		query: map[string]string{
			"query": opt.Query,
		},
	})
	if err != nil {
		return TeamPage{}, err
	}

	// Decode the response into a TeamPage response object.
	var teamPage TeamPage
	if err := decodeResponse(resp, &teamPage); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return TeamPage{}, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	return teamPage, nil
}
