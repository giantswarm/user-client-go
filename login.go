package client

import (
	"io"
	"net/http"

	"github.com/juju/errgo"
)

func (this *Client) Login(userOrMail string, reqBody io.Reader) (*http.Response, error) {
	res, err := this.post(this.endpointUrl("/user/"+userOrMail+"/login"), "application/json", reqBody)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return res, nil
}
