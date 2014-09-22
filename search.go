package client

import (
	"github.com/catalyst-zero/api-schema"
)

type SearchRequest struct {
	Username string `json:"username"`
}

func (c *Client) Search(req SearchRequest) ([]User, error) {
	var result []User

	httpResp, err := c.postJson(c.endpointUrl("/user/search"), req)
	if err != nil {
		return result, Mask(err)
	}

	// Check if request body was valid.
	if ok, err := apischema.IsStatusWrongInput(&httpResp.Body); err != nil {
		return result, Mask(err)
	} else if ok {
		return result, Mask(ErrWrongInput)
	}

	if err := apischema.ParseData(&httpResp.Body, &result); err != nil {
		return result, Mask(err)
	}
	return result, nil
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
