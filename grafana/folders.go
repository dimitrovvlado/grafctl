package grafana

import (
	"fmt"
	"net/http"
)

// ListFolders returns a list of folders.
func (c *Client) ListFolders() ([]Folder, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: FoldersEndpoint,
	})
	if err != nil {
		return nil, err
	}

	// Decode the response into a []Folder response object.
	var folders []Folder
	if err := decodeResponse(resp, &folders); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return nil, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	// Check if we didn't get a result and return an error if true.
	if folders == nil || len(folders) <= 0 {
		return make([]Folder, 0), nil
	}

	return folders, nil
}

// GetFolder returns a folder by uid
func (c *Client) GetFolder(uid string) (Folder, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("%s/%s", FoldersEndpoint, uid),
	})
	if err != nil {
		return Folder{}, err
	}

	// Decode the response into a User response object.
	var folder Folder
	if err := decodeResponse(resp, &folder); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return Folder{}, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	return folder, nil
}

//DeleteFolder deletes a single folder
func (c *Client) DeleteFolder(folder Folder) error {
	resp, err := c.doRequest(&request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("%s/%s", FoldersEndpoint, folder.UID),
	})

	defer resp.Body.Close()
	return err
}

//CreateFolder creates a single folder
func (c *Client) CreateFolder(folder Folder) (Folder, error) {
	resp, err := c.doRequest(&request{
		method:   http.MethodPost,
		endpoint: FoldersEndpoint,
		data:     folder,
	})

	var res Folder
	if err != nil {
		return res, err
	}

	if err := decodeResponse(resp, &res); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return res, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	return res, nil
}
