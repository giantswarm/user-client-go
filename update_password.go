package client

import (
	"encoding/json"
	"io"
	"strings"

	apiSchemaPkg "github.com/catalyst-zero/api-schema"
)

func (this *Client) UpdatePassword(userId, oldPass, newPass string) error {
	reader, err := newUpdatePasswordPayload(oldPass, newPass)
	if err != nil {
		return Mask(err)
	}

	res, err := this.post(this.endpointUrl("/user/"+userId+"/password/update"), "application/json", reader)
	if err != nil {
		return Mask(err)
	}

	// Check if request body was valid.
	if ok, err := apiSchemaPkg.IsStatusWrongInput(&res.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrWrongInput)
	}

	// Check if valid credentials.
	if ok, err := apiSchemaPkg.IsStatusResourceInvalidCredentials(&res.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrInvalidCredentials)
	}

	return nil
}

func newUpdatePasswordPayload(oldPass, newPass string) (io.Reader, error) {
	payload := map[string]string{
		"old_password": oldPass,
		"new_password": newPass,
	}

	byteSlice, err := json.Marshal(payload)
	if err != nil {
		return nil, Mask(err)
	}

	return strings.NewReader(string(byteSlice)), nil
}
