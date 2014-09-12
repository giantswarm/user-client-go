package client

import (
	"encoding/json"
	"io"
	"strings"

	apiSchemaPkg "github.com/catalyst-zero/api-schema"
)

func (this *Client) UpdateEmail(userId, oldMail, newMail string) error {
	reader, err := newUpdateEmailPayload(oldMail, newMail)
	if err != nil {
		return Mask(err)
	}

	res, err := this.post(this.endpointUrl("/user/"+userId+"/email/update"), "application/json", reader)
	if err != nil {
		return Mask(err)
	}

	// Check if valid credentials.
	if ok, err := apiSchemaPkg.IsStatusResourceInvalidCredentials(&res.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrInvalidCredentials)
	}

	return nil
}

func newUpdateEmailPayload(oldMail, newMail string) (io.Reader, error) {
	payload := map[string]string{
		"old_email": oldMail,
		"new_email": newMail,
	}

	byteSlice, err := json.Marshal(payload)
	if err != nil {
		return nil, Mask(err)
	}

	return strings.NewReader(string(byteSlice)), nil
}
