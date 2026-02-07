package handler

import "errors"

var (
	ErrNotFound    = errors.New("url not found")
	ErrInvalidCode = errors.New("code is invalid")
)
