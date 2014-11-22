package client

import (
	"github.com/catalyst-zero/api-schema"
)

func (this *Client) Create(username, email, password string) (string, error) {
	payload := map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	}

	resp, err := this.postSchemaJSON("/user/", payload)
	if err != nil {
		return "", Mask(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return "", Mask(err)
	}

	var userID string
	if err := resp.UnmarshalData(&userID); err != nil {
		return "", Mask(err)
	}

	if userID == "" {
		return "", Mask(ErrUnexpectedResponse)
	}
	return userID, nil
}
