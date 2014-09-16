package client

import (
	"github.com/juju/errgo"
)

var (
	ErrWrongInput         = errgo.New("Wrong input")
	ErrInvalidCredentials = errgo.New("Invalid credentials")
	ErrUnexpectedResponse = errgo.New("Unexpected response from user service")

	Mask = errgo.MaskFunc(IsErrInvalidCredentials, IsErrWrongInput)
)

func IsErrInvalidCredentials(err error) bool {
	return errgo.Cause(err) == ErrInvalidCredentials
}

func IsErrWrongInput(err error) bool {
	return errgo.Cause(err) == ErrWrongInput
}
