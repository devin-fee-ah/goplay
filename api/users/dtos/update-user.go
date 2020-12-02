package dtos

// UpdateUser dto
type UpdateUser struct {
	Age  int    `json:"age" example:"33"`
	Name string `json:"name" example:"Devin"`
}

// Validate the dto
func (dto UpdateUser) Validate() error {
	switch {
	case dto.Age < 0:
		return ErrAgeInvalid
	default:
		return nil
	}
}
