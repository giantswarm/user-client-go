package client

import (
	"github.com/juju/errgo"
)

var (
	ErrNotFound           = errgo.New("Not found")
	ErrWrongInput         = errgo.New("Wrong input")
	ErrInvalidCredentials = errgo.New("Invalid credentials")
	ErrUnexpectedResponse = errgo.New("Unexpected response from user service")

	Mask = errgo.MaskFunc(IsErrInvalidCredentials, IsErrWrongInput, IsErrNotFound)
)

func IsErrInvalidCredentials(err error) bool {
	return errgo.Cause(err) == ErrInvalidCredentials
}

func IsErrNotFound(err error) bool {
	return errgo.Cause(err) == ErrNotFound
}

func IsErrWrongInput(err error) bool {
	return errgo.Cause(err) == ErrWrongInput
}
