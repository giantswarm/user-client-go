package client

import (
	apischema "github.com/giantswarm/api-schema"
)

// UpdatePassword updates the password for the given user to newPass.
// oldPass must contain the current password, otherwise an error is thrown.
func (client *Client) UpdatePassword(userID, oldPass, newPass string) error {
	payload := map[string]string{
		"old_password": oldPass,
		"new_password": newPass,
	}
	return client.postPasswordUpdate(userID, payload)
}

// ResetPassword resets the password of the given user to the new password.
// No verification will be done.
func (client *Client) ResetPassword(userID, newPass string) error {
	payload := map[string]string{
		"new_password": newPass,
	}
	return client.postPasswordUpdate(userID, payload)
}

func (client *Client) postPasswordUpdate(userID string, payload map[string]string) error {
	resp, err := client.postSchemaJSON("/user/"+userID+"/password/update", payload)
	if err != nil {
		return Mask(err)
	}

	// Check the status is kind of expected
	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_RESOURCE_UPDATED); err != nil {
		return Mask(err)
	}
	return nil
}
