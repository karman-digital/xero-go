package xeroerrors

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

var ErrInternal = errors.New("internal error, unrelated to xero")

var ErrApiError = errors.New("xero api error")
