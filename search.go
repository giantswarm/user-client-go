package client

import (
	"github.com/catalyst-zero/api-schema"
)

type SearchRequest struct {
	Usernames []string `json:"usernames"`
	UserIDs   []string `json:"user_ids"`
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
	if err := mapCommonErrors(httpResp); err != nil {
		return nil, Mask(err)
	}

	if ok, err := apischema.IsStatusData(&httpResp.Body); err != nil {
		return nil, Mask(err)
	} else if ok {

		var result SearchResult
		if err := apischema.ParseData(&httpResp.Body, &result); err != nil {
			return nil, Mask(err)
		}
		return result.Items, nil
	}

	return nil, Mask(ErrUnexpectedResponse)
}

func (c *Client) SearchByUserIDs(userIDs []string) ([]User, error) {
	users, err := c.Search(SearchRequest{UserIDs: userIDs})
	if err != nil {
		return nil, Mask(err)
	}
	return users, nil
}

func (c *Client) SearchByUsername(username string) (User, error) {
	zeroValue := User{}

	users, err := c.Search(SearchRequest{Usernames: []string{username}})
	if err != nil {
		return zeroValue, Mask(err)
	}

	if len(users) != 1 {
		return zeroValue, Mask(ErrUnexpectedResponse)
	}

	return users[0], nil
}
