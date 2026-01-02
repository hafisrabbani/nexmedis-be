package error

import "errors"

var (
	ErrClientAlreadyExists = errors.New("client already exists")
	ErrClientNotFound      = errors.New("client not found")
	ErrInvalidAPIKey       = errors.New("invalid api key")
)
