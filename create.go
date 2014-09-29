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

	res, err := this.postJson(this.endpointUrl("/user/"), payload)
	if err != nil {
		return "", Mask(err)
	}

	if err := mapCommonErrors(res); err != nil {
		return "", Mask(err)
	}

	if ok, err := apischema.IsStatusData(&res.Body); err != nil {
		return "", Mask(err)
	} else if ok {
		var userID string
		if err := apischema.ParseData(&res.Body, &userID); err != nil {
			return "", Mask(err)
		}

		return userID, nil
	}
	return "", Mask(ErrUnexpectedResponse)
}
