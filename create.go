package client

import (
	"time"

	"github.com/giantswarm/api-schema"
)

func (this *Client) Create(username, email, password string) (string, error) {
	var noExpirationDate time.Time
	return this.CreateTimeLimited(username, email, password, noExpirationDate)
}

func (this *Client) CreateTimeLimited(username, email, password string, expirationDate time.Time) (string, error) {
	payload := map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	}

	if !expirationDate.IsZero() {
		payload["expiration_date"] = expirationDate.Format(time.RFC3339)
	}

	resp, err := this.postSchemaJSON("/user/", payload)
	if err != nil {
		return "", Mask(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return "", Mask(err)
	}

	var userID string
	if err := resp.UnmarshalData(&userID); err != nil {
		return "", Mask(err)
	}

	if userID == "" {
		return "", Mask(ErrUnexpectedResponse)
	}
	return userID, nil
}
