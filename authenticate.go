package client

import (
	"io"

	apiSchemaPkg "github.com/catalyst-zero/api-schema"
)

func (this *Client) Authenticate(userOrMail string, reqBody io.Reader) (string, error) {
	zeroVal := ""

	res, err := this.post(this.endpointUrl("/user/"+userOrMail+"/authenticate"), "application/json", reqBody)
	if err != nil {
		return zeroVal, Mask(err)
	}

	// Check user service response.
	var userId string
	if err := apiSchemaPkg.ParseData(&res.Body, &userId); err != nil {
		return zeroVal, Mask(err)
	}

	// Just proxy unexpected user service response.
	if userId == "" {
		return zeroVal, Mask(ErrUnexpectedResponse)
	}

	return userId, nil
}
