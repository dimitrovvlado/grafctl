package grafana

import (
	"fmt"
	"net/http"
)

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

// GetUser returns a user by id
func (c *Client) GetUser(ID string) (User, error) {
	endpoint := UsersEndpoint

	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: endpoint,
	})
	if err != nil {
		return User{}, err
	}

	// Decode the response into a User response object.
	var user User
	if err := decodeResponse(resp, &user); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return User{}, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	return user, nil
}
