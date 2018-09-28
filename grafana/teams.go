package grafana

import (
	"fmt"
	"net/http"
)

//GetTeam returns a team by ID
func (c *Client) GetTeam(teamID string) (Team, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s", TeamsEndpoint, teamID),
	})
	if err != nil {
		return Team{}, err
	}

	// Decode the response into a TeamPage response object.
	var team Team
	if err := decodeResponse(resp, &team); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return Team{}, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	return team, nil
}

//SearchTeams returns a list of teams
func (c *Client) SearchTeams(opt *SearchTeamsOptions) (TeamPage, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: TeamsSearchEndpoint,
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

//CreateTeam creates a team
func (c *Client) CreateTeam(team Team) (int, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodPost,
		endpoint: TeamsEndpoint,
		data:     team,
	})
	if err != nil {
		return -1, err
	}

	// Decode the response into a TeamPage response object.
	var teamResponse map[string]interface{}
	if err := decodeResponse(resp, &teamResponse); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return -1, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}
	return int(teamResponse["teamId"].(float64)), nil
}

// DeleteTeam deletes a team
func (c *Client) DeleteTeam(team Team) error {
	resp, err := c.doRequest(&request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("%s/%d", TeamsEndpoint, team.ID),
	})
	defer resp.Body.Close()
	return err
}
