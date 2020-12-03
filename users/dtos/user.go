package dtos

// User example
type User struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"user name"`
	Age  int    `json:"uuid" example:"25" format:"int64"`
}
