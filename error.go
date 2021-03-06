package client

import (
	"strings"

	"github.com/giantswarm/api-schema"
	"github.com/juju/errgo"
)

const (
	reasonEmailMustBeVerified = "email_must_be_verified" // This string is generated in userd to identify these errors so keep it constant
	reasonUserExpired         = "user_expired"
)

var (
	ErrNotFound                = errgo.New("Not found")
	ErrWrongInput              = errgo.New("Wrong input")
	ErrInvalidCredentials      = errgo.New("Invalid credentials")
	ErrUnexpectedResponse      = errgo.New("Unexpected response from user service")
	ErrUserAlreadyExists       = errgo.New("User with username or email already exists.")
	ErrUserEmailMustBeVerified = errgo.New("Email must be verified to authenticate.")

	ErrUserExpired = errgo.New(reasonUserExpired)

	Mask = errgo.MaskFunc(IsErrInvalidCredentials, IsErrUserEmailMustBeVerified, IsErrWrongInput, IsErrNotFound)
)

func IsErrInvalidCredentials(err error) bool {
	return errgo.Cause(err) == ErrInvalidCredentials || apischema.IsResourceInvalidCredentialsError(errgo.Cause(err))
}

func IsErrNotFound(err error) bool {
	return errgo.Cause(err) == ErrNotFound || apischema.IsResourceNotFoundError(errgo.Cause(err))
}

// IsErrUserEmailMustBeVerified checks if the given error is a ErrUserEmailMustBeVerified.
// Note that this is a special case of ErrWrongInput so check this first.
func IsErrUserEmailMustBeVerified(err error) bool {
	cause := errgo.Cause(err)
	if cause == ErrUserEmailMustBeVerified {
		return true
	}
	return apischema.IsWrongInputError(cause) && strings.Contains(cause.Error(), reasonEmailMustBeVerified)
}

// IsErrUserExpired returns true if the given error was caused by ErrUserExpired
// or is an apischema error for an 'user_expired'.
func IsErrUserExpired(err error) bool {
	cause := errgo.Cause(err)
	if cause == ErrUserExpired {
		return true
	}
	return apischema.IsWrongInputError(cause) && strings.Contains(cause.Error(), reasonUserExpired)
}

func IsErrWrongInput(err error) bool {
	return errgo.Cause(err) == ErrWrongInput || apischema.IsWrongInputError(errgo.Cause(err))
}

func IsErrUserAlreadyExists(err error) bool {
	return errgo.Cause(err) == ErrUserAlreadyExists || apischema.IsResourceAlreadyExistsError(errgo.Cause(err))
}
