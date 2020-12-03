package dtos

import "fmt"

// AddUser example
type AddUser struct {
	Age  int    `json:"age" example:"33" binding:"required"`
	Name string `json:"name" example:"Devin" binding:"required"`
}

// Validate example
func (dto AddUser) Validate() error {
	switch {
	case len(dto.Name) == 0:
		return ErrNameEmpty
	case dto.Age <= 0:
		fmt.Printf("******type is %T\n", dto.Age)
		fmt.Printf("******value is %d\n", dto.Age)
		return ErrAgeInvalid

	default:
		return nil
	}
}
