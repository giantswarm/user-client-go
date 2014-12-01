package client

import (
	"github.com/catalyst-zero/api-schema"
	"github.com/juju/errgo"
)

var (
	ErrNotFound           = errgo.New("Not found")
	ErrWrongInput         = errgo.New("Wrong input")
	ErrInvalidCredentials = errgo.New("Invalid credentials")
	ErrUnexpectedResponse = errgo.New("Unexpected response from user service")
	ErrUserAlreadyExists  = errgo.New("User with username or email already exists.")

	Mask = errgo.MaskFunc(IsErrInvalidCredentials, IsErrWrongInput, IsErrNotFound)
)

func IsErrInvalidCredentials(err error) bool {
	return errgo.Cause(err) == ErrInvalidCredentials || apischema.IsResourceInvalidCredentialsError(errgo.Cause(err))
}

func IsErrNotFound(err error) bool {
	return errgo.Cause(err) == ErrNotFound || apischema.IsResourceNotFoundError(errgo.Cause(err))
}

func IsErrWrongInput(err error) bool {
	return errgo.Cause(err) == ErrWrongInput || apischema.IsWrongInputError(errgo.Cause(err))
}

func IsErrUserAlreadyExists(err error) bool {
	return errgo.Cause(err) == ErrUserAlreadyExists || apischema.IsResourceAlreadyExists(errgo.Cause(err))
}
