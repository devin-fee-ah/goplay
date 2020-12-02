package dtos

// UpdateUser example
type UpdateUser struct {
	Name string `json:"name" example:"user name"`
}

// Validate example
func (a UpdateUser) Validate() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}
