package models

import (
	"errors"
)

// User example
type User struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"user name"`
	Age  int    `json:"uuid" example:"25" format:"int64"`
}

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

// // UsersAll example
// func UsersAll(q string) ([]User, error) {
// 	if q == "" {
// 		return users, nil
// 	}
// 	as := []User{}
// 	for k, v := range users {
// 		if q == v.Name {
// 			as = append(as, users[k])
// 		}
// 	}
// 	return as, nil
// }

// // UserOne example
// func UserOne(id int) (User, error) {
// 	for _, v := range users {
// 		if id == v.ID {
// 			return v, nil
// 		}
// 	}
// 	return User{}, ErrNoRow
// }

// // Insert example
// func (a User) Insert() (int, error) {
// 	userMaxID++
// 	a.ID = userMaxID
// 	a.Name = fmt.Sprintf("user_%d", userMaxID)
// 	users = append(users, a)
// 	return userMaxID, nil
// }

// // Delete example
// func Delete(id int) error {
// 	for k, v := range users {
// 		if id == v.ID {
// 			users = append(users[:k], users[k+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("user id=%d is not found", id)
// }

// // Update example
// func (a User) Update() error {
// 	for k, v := range users {
// 		if a.ID == v.ID {
// 			users[k].Name = a.Name
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("user id=%d is not found", a.ID)
// }
