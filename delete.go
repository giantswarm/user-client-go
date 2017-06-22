package client

import (
	"github.com/giantswarm/api-schema"
)

func (this *Client) Delete(userID string) (ok string, err error) {
	resp, err = this.postSchema("/user/" + userID + "/delete")
	if err != nil {
		return Mask(err)
	}

	// Check the status is kind of expected
	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_RESOURCE_DELETED); err != nil {
		return Mask(err)
	}
	return nil
}
