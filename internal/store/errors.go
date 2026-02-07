package store

import "errors"

var (
	ErrNotFound         = errors.New("url not found")
	ErrAlreadyExists    = errors.New("url already exists")
	ErrStoreUnavailable = errors.New("store unavailable")
)
