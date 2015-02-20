package client

import (
	"github.com/giantswarm/api-schema"
)

func (this *Client) UpdateEmail(userId, oldMail, newMail string) error {
	payload := map[string]string{
		"old_email": oldMail,
		"new_email": newMail,
	}

	resp, err := this.postSchemaJSON("/user/"+userId+"/email/update", payload)
	if err != nil {
		return Mask(err)
	}

	// Check the status is kind of expected
	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_RESOURCE_UPDATED); err != nil {
		return Mask(err)
	}
	return nil
}
