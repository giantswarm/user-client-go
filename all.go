package client

import "github.com/giantswarm/api-schema"

func (this *Client) All() ([]User, error) {
	emptyResult := []User{}

	resp, err := apischema.FromHTTPResponse(this.get(this.endpointUrl("/users/")))
	if err != nil {
		return emptyResult, err
	}

	// Check the status is kind of expected
	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return emptyResult, Mask(err)
	}

	var searchResult SearchResult
	if err := resp.UnmarshalData(&searchResult); err != nil {
		return emptyResult, Mask(err)
	}

	return searchResult.Items, nil
}
