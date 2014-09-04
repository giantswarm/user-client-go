package client

import (
	"github.com/juju/errgo"
)

var (
	ErrInvalidCredentials = errgo.New("Invalid credentials")
	ErrUnexpectedResponse = errgo.New("Unexpected response from user service")

	Mask = errgo.MaskFunc(IsErrInvalidCredentials)
)

func IsErrInvalidCredentials(err error) bool {
	return errgo.Cause(err) == ErrInvalidCredentials
}
