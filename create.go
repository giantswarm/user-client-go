package client

import (
	"github.com/catalyst-zero/api-schema"
	"github.com/juju/errgo"
)

func (this *Client) Create(username, email, password string) (string, error) {
	payload := map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	}

	res, err := this.postJson(this.endpointUrl("/user/"), payload)
	if err != nil {
		return "", errgo.Mask(err)
	}

	var userID string

	if ok, err := apischema.IsStatusData(&res.Body); err != nil {
		return "", errgo.Mask(err)
	} else if ok {
		if err := apischema.ParseData(&res.Body, &userID); err != nil {
			return "", errgo.Mask(err)
		}
	}

	return userID, nil
}
