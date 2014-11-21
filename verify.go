package client

import (
	"github.com/catalyst-zero/api-schema"
)

func (this *Client) SetEmailVerified(userId string) error {
	resp, err := this.postSchema("/user/" + userId + "/email/verify")
	if err != nil {
		return Mask(err)
	}

	// Check the status is kind of expected
	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_RESOURCE_UPDATED); err != nil {
		return Mask(err)
	}
	return nil
}
