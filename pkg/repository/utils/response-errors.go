package utils

import "errors"

type DbResponseCode int

var (
	ErrAlreadyExists  = errors.New("similar record already exists")
	ErrDbNotConnected = errors.New("")
)
