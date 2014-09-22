package client

import (
	"encoding/json"
)

type SearchRequest struct {
	Username string `json:"username"`
}

func (c *Client) Search(req SearchRequest) ([]User, error) {
	var result []User

	httpResp, err := c.postJson(this.endpointUrl("/user/search"), req)
	if err != nil {
		return result, Mask(err)
	}

	if err := json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
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
