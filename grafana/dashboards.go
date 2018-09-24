package grafana

import (
	"fmt"
	"net/http"
)

// ListDashboards returns a list of dashboards.
func (c *Client) ListDashboards() ([]Dashboard, error) {

	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: SearchEndpoint,
	})
	if err != nil {
		return nil, err
	}

	var unf []Dashboard
	if err := decodeResponse(resp, &unf); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return nil, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}
	dashboards := unf[:0]
	for _, u := range unf {
		if u.Type == "dash-db" {
			dashboards = append(dashboards, u)
		}
	}

	// Check if we didn't get a result and return an error if true.
	if dashboards == nil || len(dashboards) <= 0 {
		return make([]Dashboard, 0), nil
	}

	return dashboards, nil
}

// CreateDashboard creates a dashboard.
func (c *Client) CreateDashboard(db DashboardRequest) (DashboardResponse, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodPost,
		endpoint: DashboardsImportEndpoint,
		data:     db,
	})
	if err != nil {
		return DashboardResponse{}, err
	}

	var dr DashboardResponse
	if err := decodeResponse(resp, &dr); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return DashboardResponse{}, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	return dr, nil
}