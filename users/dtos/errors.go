package dtos

import "errors"

var (
	// ErrAgeInvalid when an age is invalid
	ErrAgeInvalid = errors.New("age must be > 0")
	// ErrNameEmpty when a name is emmpty
	ErrNameEmpty = errors.New("name is empty")
)
