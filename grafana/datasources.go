package grafana

import (
	"fmt"
	"net/http"
)

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
