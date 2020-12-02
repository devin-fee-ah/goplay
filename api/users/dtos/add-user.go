package dtos

import (
	"errors"
)

//  example
var (
	ErrNameInvalid = errors.New("name is empty")
)

// AddUser example
type AddUser struct {
	Name string `json:"name" example:"user name"`
}

// Validate example
func (a AddUser) Validate() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}
