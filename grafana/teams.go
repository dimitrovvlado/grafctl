package grafana

import (
	"fmt"
	"net/http"
)

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
