package client

import (
	"net/http"

	"github.com/juju/errgo"
)

func (this *Client) Create(username, email, password string) (*http.Response, error) {
	payload := map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	}

	res, err := this.postJson(this.endpointUrl("/user/"), payload)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return res, nil
}
