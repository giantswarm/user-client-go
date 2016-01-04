package client

import (
	"time"

	"github.com/giantswarm/api-schema"
)

func (this *Client) UpdateExpirationDate(userId string, expirationDate time.Time) error {
	payload := map[string]string{}

	if !expirationDate.IsZero() {
		payload["expiration_date"] = expirationDate.Format(time.RFC3339)
	}

	resp, err := this.postSchemaJSON("/user/"+userId+"/expiration_date/update", payload)
	if err != nil {
		return Mask(err)
	}

	// Check the status is kind of expected
	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_RESOURCE_UPDATED); err != nil {
		return Mask(err)
	}
	return nil
}
