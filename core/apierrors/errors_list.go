package apierrors

import "errors"

var (
	// ErrNotFound is not found err
	ErrNotFound     = errors.New("object not found")
	ErrInternal     = errors.New("internal server error")
	ErrNoAuthHeader = errors.New("authorization header not provided")
	ErrInvalidAuth  = errors.New("invalid authorization header")
	ErrInvalidLogin = errors.New("invalid credentials provided")
)
