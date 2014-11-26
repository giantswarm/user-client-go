package client

import (
	"github.com/catalyst-zero/api-schema"
)

type SearchRequest struct {
	Usernames []string `json:"usernames"`
	Emails    []string `json:"emails"`
	UserIDs   []string `json:"user_ids"`
}

type SearchResult struct {
	Size  int    `json:"size"`
	Items []User `json:"items"`
}

func (c *Client) Search(req SearchRequest) (SearchResult, error) {
	zeroVal := SearchResult{}

	resp, err := c.postSchemaJSON("/user/search", req)
	if err != nil {
		return zeroVal, Mask(err)
	}

	// Check the status is kind of expected
	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return zeroVal, Mask(err)
	}

	var result SearchResult
	if err := resp.UnmarshalData(&result); err != nil {
		return zeroVal, Mask(err)
	}

	return result, nil
}

func (c *Client) SearchByUserIDs(userIDs []string) ([]User, error) {
	result, err := c.Search(SearchRequest{UserIDs: userIDs})
	if err != nil {
		return nil, Mask(err)
	}
	return result.Items, nil
}

func (c *Client) SearchByUsername(username string) (User, error) {
	zeroValue := User{}

	result, err := c.Search(SearchRequest{Usernames: []string{username}})
	if err != nil {
		return zeroValue, Mask(err)
	}

	if len(result.Items) == 0 {
		return zeroValue, Mask(ErrNotFound)
	}
	if len(result.Items) > 1 {
		return zeroValue, Mask(ErrUnexpectedResponse)
	}

	return result.Items[0], nil
}

func (c *Client) SearchByEmail(email string) (User, error) {
	zeroValue := User{}

	result, err := c.Search(SearchRequest{Emails: []string{email}})
	if err != nil {
		return zeroValue, Mask(err)
	}

	if len(result.Items) == 0 {
		return zeroValue, Mask(ErrNotFound)
	}
	if len(result.Items) > 1 {
		return zeroValue, Mask(ErrUnexpectedResponse)
	}

	return result.Items[0], nil
}
