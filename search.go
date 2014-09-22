package client

import (
	"github.com/catalyst-zero/api-schema"
)

type SearchRequest struct {
	Username string `json:"username"`
}

type SearchResult struct {
	Size  int `json:"size"`
	Items []User
}

func (c *Client) Search(req SearchRequest) ([]User, error) {
	httpResp, err := c.postJson(c.endpointUrl("/user/search"), req)
	if err != nil {
		return nil, Mask(err)
	}

	// Check if request body was valid.
	if ok, err := apischema.IsStatusWrongInput(&httpResp.Body); err != nil {
		return nil, Mask(err)
	} else if ok {
		return nil, Mask(ErrWrongInput)
	}

	var result SearchResult
	if err := apischema.ParseData(&httpResp.Body, &result); err != nil {
		return nil, Mask(err)
	}
	return result.Items, nil
}

func (c *Client) SearchByUsername(username string) (User, error) {
	zeroValue := User{}

	users, err := c.Search(SearchRequest{Username: username})
	if err != nil {
		return zeroValue, Mask(err)
	}

	if len(users) != 1 {
		return zeroValue, Mask(ErrUnexpectedResponse)
	}

	return users[0], nil
}
