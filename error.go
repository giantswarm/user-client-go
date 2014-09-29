package client

import (
	"github.com/catalyst-zero/api-schema"
	"github.com/juju/errgo"

	"net/http"
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

func mapCommonErrors(response *http.Response) error {
	if ok, err := apischema.IsStatusWrongInput(&response.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrWrongInput)
	}

	if ok, err := apischema.IsStatusResourceNotFound(&response.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrNotFound)
	}

	if ok, err := apischema.IsStatusResourceInvalidCredentials(&response.Body); err != nil {
		return Mask(err)
	} else if ok {
		return Mask(ErrInvalidCredentials)
	}

	return nil
}
