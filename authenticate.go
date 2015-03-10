package client

import (
	"github.com/giantswarm/api-schema"
)

// Authenticate checks that a user with the given username (or email) exists.
func (this *Client) AuthenticateCredentials(userOrMail, password string) (string, error) {
	zeroVal := ""

	payload := map[string]string{
		"password": password,
	}

	resp, err := this.postSchemaJSON("/user/"+userOrMail+"/authenticate", payload)
	if err != nil {
		return zeroVal, Mask(err)
	}

	// Check the status is kind of expected
	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return zeroVal, Mask(err)
	}

	// Check user service response.
	var userId string
	if err := resp.UnmarshalData(&userId); err != nil {
		return zeroVal, Mask(err)
	}

	// Just proxy unexpected user service response.
	if userId == "" {
		return zeroVal, Mask(ErrUnexpectedResponse)
	}

	return userId, nil
}
