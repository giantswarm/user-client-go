package client

import (
	"github.com/catalyst-zero/api-schema"
)

func (this *Client) UpdatePassword(userId, oldPass, newPass string) error {
	payload := map[string]string{
		"old_password": oldPass,
		"new_password": newPass,
	}

	resp, err := this.postSchemaJSON("/user/"+userId+"/password/update", payload)
	if err != nil {
		return Mask(err)
	}

	// Check the status is kind of expected
	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_RESOURCE_UPDATED); err != nil {
		return Mask(err)
	}
	return nil
}
