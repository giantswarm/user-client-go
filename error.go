package client

import (
	"github.com/juju/errgo"
)

var (
	ErrUnexpectedResponse = errgo.New("Unexpected response from user service")

	Mask = errgo.MaskFunc()
)
